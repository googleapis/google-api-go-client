// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

type unexpectedReader struct{}

func (unexpectedReader) Read([]byte) (int, error) {
	return 0, fmt.Errorf("unexpected read in test")
}

// event is an expected request/response pair
type event struct {
	// the byte range header that should be present in a request.
	byteRange string
	// the http status code to send in response.
	responseStatus int
}

// interruptibleTransport is configured with a canned set of requests/responses.
// It records the incoming data, unless the corresponding event is configured to return
// http.StatusServiceUnavailable.
type interruptibleTransport struct {
	events []event
	buf    []byte
	bodies bodyTracker
}

// bodyTracker keeps track of response bodies that have not been closed.
type bodyTracker map[io.ReadCloser]struct{}

func (bt bodyTracker) Add(body io.ReadCloser) {
	bt[body] = struct{}{}
}

func (bt bodyTracker) Close(body io.ReadCloser) {
	delete(bt, body)
}

type trackingCloser struct {
	io.Reader
	tracker bodyTracker
}

func (tc *trackingCloser) Close() error {
	tc.tracker.Close(tc)
	return nil
}

func (tc *trackingCloser) Open() {
	tc.tracker.Add(tc)
}

func (t *interruptibleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(t.events) == 0 {
		panic("ran out of events, but got a request")
	}
	ev := t.events[0]
	t.events = t.events[1:]
	if got, want := req.Header.Get("Content-Range"), ev.byteRange; got != want {
		return nil, fmt.Errorf("byte range: got %s; want %s", got, want)
	}

	if ev.responseStatus != http.StatusServiceUnavailable {
		buf, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading from request data: %v", err)
		}
		t.buf = append(t.buf, buf...)
	}

	tc := &trackingCloser{unexpectedReader{}, t.bodies}
	tc.Open()
	h := http.Header{}
	status := ev.responseStatus

	// Support "X-GUploader-No-308" like Google:
	if status == 308 && req.Header.Get("X-GUploader-No-308") == "yes" {
		status = 200
		h.Set("X-Http-Status-Code-Override", "308")
	}

	res := &http.Response{
		StatusCode: status,
		Header:     h,
		Body:       tc,
	}
	return res, nil
}

// progressRecorder records updates, and calls f for every invocation of ProgressUpdate.
type progressRecorder struct {
	updates []int64
	f       func()
}

func (pr *progressRecorder) ProgressUpdate(current int64) {
	pr.updates = append(pr.updates, current)
	if pr.f != nil {
		pr.f()
	}
}

func TestInterruptedTransferChunks(t *testing.T) {
	type testCase struct {
		name         string
		data         string
		chunkSize    int
		events       []event
		wantProgress []int64
	}

	for _, tc := range []testCase{
		{
			name:      "large",
			data:      strings.Repeat("a", 300),
			chunkSize: 90,
			events: []event{
				{"bytes 0-89/*", http.StatusServiceUnavailable},
				{"bytes 0-89/*", 308},
				{"bytes 90-179/*", 308},
				{"bytes 180-269/*", http.StatusServiceUnavailable},
				{"bytes 180-269/*", 308},
				{"bytes 270-299/300", 200},
			},
			wantProgress: []int64{90, 180, 270, 300},
		},
		{
			name:      "small",
			data:      strings.Repeat("a", 20),
			chunkSize: 10,
			events: []event{
				{"bytes 0-9/*", http.StatusServiceUnavailable},
				{"bytes 0-9/*", 308},
				{"bytes 10-19/*", http.StatusServiceUnavailable},
				{"bytes 10-19/*", 308},
				// 0 byte final request demands a byte range with leading asterix.
				{"bytes */20", http.StatusServiceUnavailable},
				{"bytes */20", 200},
			},
			wantProgress: []int64{10, 20},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			media := strings.NewReader(tc.data)

			tr := &interruptibleTransport{
				buf:    make([]byte, 0, len(tc.data)),
				events: tc.events,
				bodies: bodyTracker{},
			}

			pr := progressRecorder{}
			rx := &ResumableUpload{
				Client:    &http.Client{Transport: tr},
				Media:     NewMediaBuffer(media, tc.chunkSize),
				MediaType: "text/plain",
				Callback:  pr.ProgressUpdate,
			}

			oldBackoff := backoff
			backoff = func() Backoff { return new(NoPauseBackoff) }
			defer func() { backoff = oldBackoff }()

			res, err := rx.Upload(context.Background())
			if err == nil {
				res.Body.Close()
			}
			if err != nil || res == nil || res.StatusCode != http.StatusOK {
				if res == nil {
					t.Fatalf("Upload not successful, res=nil: %v", err)
				} else {
					t.Fatalf("Upload not successful, statusCode=%v, err=%v", res.StatusCode, err)
				}
			}
			if !reflect.DeepEqual(tr.buf, []byte(tc.data)) {
				t.Fatalf("transferred contents:\ngot %s\nwant %s", tr.buf, tc.data)
			}

			if !reflect.DeepEqual(pr.updates, tc.wantProgress) {
				t.Fatalf("progress updates: got %v, want %v", pr.updates, tc.wantProgress)
			}

			if len(tr.events) > 0 {
				t.Fatalf("did not observe all expected events.  leftover events: %v", tr.events)
			}
			if len(tr.bodies) > 0 {
				t.Errorf("unclosed request bodies: %v", tr.bodies)
			}
		})
	}
}

func TestCancelUploadFast(t *testing.T) {
	const (
		chunkSize = 90
		mediaSize = 300
	)
	media := strings.NewReader(strings.Repeat("a", mediaSize))

	tr := &interruptibleTransport{
		buf: make([]byte, 0, mediaSize),
	}

	pr := progressRecorder{}
	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewMediaBuffer(media, chunkSize),
		MediaType: "text/plain",
		Callback:  pr.ProgressUpdate,
	}

	oldBackoff := backoff
	backoff = func() Backoff { return new(NoPauseBackoff) }
	defer func() { backoff = oldBackoff }()

	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc() // stop the upload that hasn't started yet
	res, err := rx.Upload(ctx)
	if err != context.Canceled {
		t.Fatalf("Upload err: got: %v; want: context cancelled", err)
	}
	if res != nil {
		t.Fatalf("Upload result: got: %v; want: nil", res)
	}
	if pr.updates != nil {
		t.Errorf("progress updates: got %v; want: nil", pr.updates)
	}
}

func TestCancelUploadBasic(t *testing.T) {
	const (
		chunkSize = 90
		mediaSize = 300
	)
	media := strings.NewReader(strings.Repeat("a", mediaSize))

	tr := &interruptibleTransport{
		buf: make([]byte, 0, mediaSize),
		events: []event{
			{"bytes 0-89/*", http.StatusServiceUnavailable},
			{"bytes 0-89/*", 308},
			{"bytes 90-179/*", 308},
			{"bytes 180-269/*", 308}, // Upload should be cancelled before this event.
		},
		bodies: bodyTracker{},
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	numUpdates := 0

	pr := progressRecorder{f: func() {
		numUpdates++
		if numUpdates >= 2 {
			cancelFunc()
		}
	}}

	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewMediaBuffer(media, chunkSize),
		MediaType: "text/plain",
		Callback:  pr.ProgressUpdate,
	}

	oldBackoff := backoff
	backoff = func() Backoff { return new(PauseOneSecond) }
	defer func() { backoff = oldBackoff }()

	res, err := rx.Upload(ctx)
	if err != context.Canceled {
		t.Fatalf("Upload err: got: %v; want: context cancelled", err)
	}
	if res != nil {
		t.Fatalf("Upload result: got: %v; want: nil", res)
	}
	if got, want := tr.buf, []byte(strings.Repeat("a", chunkSize*2)); !reflect.DeepEqual(got, want) {
		t.Fatalf("transferred contents:\ngot %s\nwant %s", got, want)
	}
	if got, want := pr.updates, []int64{chunkSize, chunkSize * 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("progress updates: got %v; want: %v", got, want)
	}
	if len(tr.bodies) > 0 {
		t.Errorf("unclosed request bodies: %v", tr.bodies)
	}
}

func TestRetry_EachChunkHasItsOwnRetryDeadline(t *testing.T) {
	const (
		chunkSize = 90
		mediaSize = 300
	)
	media := strings.NewReader(strings.Repeat("a", mediaSize))

	// This transport returns multiple errors on both the first chunk and third
	// chunk of the upload. If the timeout were not reset between chunks, the
	// errors on the third chunk would not retry and cause a failure.
	tr := &interruptibleTransport{
		buf: make([]byte, 0, mediaSize),
		events: []event{
			{"bytes 0-89/*", http.StatusServiceUnavailable},
			// cum: 1s sleep
			{"bytes 0-89/*", http.StatusServiceUnavailable},
			// cum: 2s sleep
			{"bytes 0-89/*", http.StatusServiceUnavailable},
			// cum: 3s sleep
			{"bytes 0-89/*", http.StatusServiceUnavailable},
			// cum: 4s sleep
			{"bytes 0-89/*", 308},
			// cum: 1s sleep <-- resets because it's a new chunk
			{"bytes 90-179/*", 308},
			// cum: 1s sleep <-- resets because it's a new chunk
			{"bytes 180-269/*", http.StatusServiceUnavailable},
			// cum: 1s sleep on later chunk
			{"bytes 180-269/*", http.StatusServiceUnavailable},
			// cum: 2s sleep on later chunk
			{"bytes 180-269/*", 308},
			// cum: 3s sleep <-- resets because it's a new chunk
			{"bytes 270-299/300", 200},
		},
		bodies: bodyTracker{},
	}

	rx := &ResumableUpload{
		Client:             &http.Client{Transport: tr},
		Media:              NewMediaBuffer(media, chunkSize),
		MediaType:          "text/plain",
		Callback:           func(int64) {},
		ChunkRetryDeadline: 5 * time.Second,
	}

	oldBackoff := backoff
	backoff = func() Backoff { return new(PauseOneSecond) }
	defer func() { backoff = oldBackoff }()

	resCode := make(chan int, 1)
	go func() {
		resp, err := rx.Upload(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		resCode <- resp.StatusCode
	}()

	select {
	case <-time.After(15 * time.Second):
		t.Fatal("timed out waiting for Upload to complete")
	case got := <-resCode:
		if want := http.StatusOK; got != want {
			t.Fatalf("want %d, got %d", want, got)
		}
	}
}
