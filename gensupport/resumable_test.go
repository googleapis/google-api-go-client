// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
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
// it records the incoming data, unless the corresponding event is configured to return
// http.StatusServiceUnavailable.
type interruptibleTransport struct {
	events []event
	buf    []byte
}

func (t *interruptibleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
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

	f := ioutil.NopCloser(unexpectedReader{})
	res := &http.Response{
		StatusCode: ev.responseStatus,
		Header:     http.Header{},
		Body:       f,
	}
	return res, nil
}

type progressRecorder struct {
	updates []int64
}

func (pr *progressRecorder) ProgressUpdate(current int64) {
	pr.updates = append(pr.updates, current)
}

func TestInterruptedTransferChunks(t *testing.T) {
	type testCase struct {
		data         string
		chunkSize    int
		events       []event
		wantProgress []int64
	}

	for _, tc := range []testCase{
		{
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
		media := strings.NewReader(tc.data)

		tr := &interruptibleTransport{
			buf:    make([]byte, 0, len(tc.data)),
			events: tc.events,
		}

		// TODO(mcgreevy): replace this sleep with something cleaner.
		sleep = func(time.Duration) {} // override time.Sleep
		pr := progressRecorder{}
		rx := &ResumableUpload{
			Client:    &http.Client{Transport: tr},
			Media:     NewResumableBuffer(media, tc.chunkSize),
			MediaType: "text/plain",
			Callback:  pr.ProgressUpdate,
		}
		res, err := rx.Upload(context.Background())
		if err != nil || res == nil || res.StatusCode != http.StatusOK {
			if res == nil {
				t.Errorf("Upload not successful, res=nil: %v", err)
			} else {
				t.Errorf("Upload not successful, statusCode=%v: %v", res.StatusCode, err)
			}
		}
		if !reflect.DeepEqual(tr.buf, []byte(tc.data)) {
			t.Errorf("transferred contents:\ngot %s\nwant %s", tr.buf, tc.data)
		}

		if !reflect.DeepEqual(pr.updates, tc.wantProgress) {
			t.Errorf("progress updates: got %v, want %v", pr.updates, tc.wantProgress)
		}

		if len(tr.events) > 0 {
			t.Errorf("did not observe all expected events.  leftover events: %v", tr.events)
		}
	}
}

func TestCancelUpload(t *testing.T) {
	chunkSize := 90
	mediaSize := 300
	media := strings.NewReader(strings.Repeat("a", mediaSize))

	tr := &interruptibleTransport{
		buf: make([]byte, 0, mediaSize),
	}

	// TODO(mcgreevy): replace this sleep with something cleaner.
	// At that time, test cancelling upload at some point other than before it starts.
	sleep = func(time.Duration) {} // override time.Sleep
	pr := progressRecorder{}
	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewResumableBuffer(media, chunkSize),
		MediaType: "text/plain",
		Callback:  pr.ProgressUpdate,
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc() // stop the upload that hasn't started yet
	res, err := rx.Upload(ctx)
	if err != context.Canceled {
		t.Errorf("Upload err: got: %v; want: context cancelled", err)
	}
	if res != nil {
		t.Errorf("Upload result: got: %v; want: nil", res)
	}
	if pr.updates != nil {
		t.Errorf("progress updates: got %v; want: nil", pr.updates)
	}
}
