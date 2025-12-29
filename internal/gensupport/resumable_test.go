// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
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
	// delay to simulate network latency for this specific event.
	delay time.Duration
}

// interruptibleTransport is configured with a canned set of requests/responses.
// It records the incoming data, unless the corresponding event is configured to return
// http.StatusServiceUnavailable.
type interruptibleTransport struct {
	events      []event
	buf         []byte
	bodies      bodyTracker
	finalHeader http.Header
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
	stallTimer := time.NewTimer(ev.delay)
	select {
	case <-stallTimer.C:
	case <-req.Context().Done():
		return nil, req.Context().Err()
	}
	if got, want := req.Header.Get("Content-Range"), ev.byteRange; got != want {
		return nil, fmt.Errorf("byte range: got %s; want %s", got, want)
	}
	t.finalHeader = req.Header

	if ev.responseStatus != http.StatusServiceUnavailable {
		buf, err := io.ReadAll(req.Body)
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
				{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
				{byteRange: "bytes 0-89/*", responseStatus: 308},
				{byteRange: "bytes 90-179/*", responseStatus: 308},
				{byteRange: "bytes 180-269/*", responseStatus: http.StatusServiceUnavailable},
				{byteRange: "bytes 180-269/*", responseStatus: 308},
				{byteRange: "bytes 270-299/300", responseStatus: 200},
			},
			wantProgress: []int64{90, 180, 270, 300},
		},
		{
			name:      "small",
			data:      strings.Repeat("a", 20),
			chunkSize: 10,
			events: []event{
				{byteRange: "bytes 0-9/*", responseStatus: http.StatusServiceUnavailable},
				{byteRange: "bytes 0-9/*", responseStatus: 308},
				{byteRange: "bytes 10-19/*", responseStatus: http.StatusServiceUnavailable},
				{byteRange: "bytes 10-19/*", responseStatus: 308},
				// 0 byte final request demands a byte range with leading asterix.
				{byteRange: "bytes */20", responseStatus: http.StatusServiceUnavailable},
				{byteRange: "bytes */20", responseStatus: 200},
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
				Media:     NewMediaBuffer(media, tc.chunkSize, false),
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
		Media:     NewMediaBuffer(media, chunkSize, false),
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
			{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
			{byteRange: "bytes 0-89/*", responseStatus: 308},
			{byteRange: "bytes 90-179/*", responseStatus: 308},
			{byteRange: "bytes 180-269/*", responseStatus: 308}, // Upload should be cancelled before this event.
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
		Media:     NewMediaBuffer(media, chunkSize, false),
		MediaType: "text/plain",
		Callback:  pr.ProgressUpdate,
	}

	oldBackoff := backoff
	backoff = func() Backoff { return new(PauseOneSecond) }
	defer func() { backoff = oldBackoff }()

	res, err := rx.Upload(ctx)
	if !errors.Is(err, context.Canceled) {
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
			{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 1s sleep
			{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 2s sleep
			{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 3s sleep
			{byteRange: "bytes 0-89/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 4s sleep
			{byteRange: "bytes 0-89/*", responseStatus: 308},
			// cum: 1s sleep <-- resets because it's a new chunk
			{byteRange: "bytes 90-179/*", responseStatus: 308},
			// cum: 1s sleep <-- resets because it's a new chunk
			{byteRange: "bytes 180-269/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 1s sleep on later chunk
			{byteRange: "bytes 180-269/*", responseStatus: http.StatusServiceUnavailable},
			// cum: 2s sleep on later chunk
			{byteRange: "bytes 180-269/*", responseStatus: 308},
			// cum: 3s sleep <-- resets because it's a new chunk
			{byteRange: "bytes 270-299/300", responseStatus: 200},
		},
		bodies: bodyTracker{},
	}

	rx := &ResumableUpload{
		Client:             &http.Client{Transport: tr},
		Media:              NewMediaBuffer(media, chunkSize, false),
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

// slowMediaReader wraps an io.Reader, introducing a delay on each Read call.
type slowMediaReader struct {
	r     io.Reader
	delay time.Duration
}

func (s *slowMediaReader) Read(p []byte) (n int, err error) {
	time.Sleep(s.delay)
	return s.r.Read(p)
}

func TestChunkTransferTimeout(t *testing.T) {
	const (
		data = "some media data" // length is 15
	)
	tests := []struct {
		name                 string
		mediaReadDelay       time.Duration
		chunkTransferTimeout time.Duration
		events               []event
		shouldFail           bool
		wantError            error
	}{
		{
			name:                 "media-read-delay-chunk-upload-succeeds",
			mediaReadDelay:       120 * time.Millisecond,
			chunkTransferTimeout: 30 * time.Millisecond,
			events: []event{
				// When media size is a multiple of chunk size, a non-final
				// chunk is sent, followed by a final, zero-byte chunk.
				{byteRange: "bytes 0-14/*", responseStatus: 308},
				{byteRange: "bytes */15", responseStatus: http.StatusOK},
			},
			shouldFail: false,
			wantError:  nil,
		},
		{
			name:                 "network-delay-chunk-upload-fails",
			mediaReadDelay:       0,
			chunkTransferTimeout: 30 * time.Millisecond,
			events: []event{
				// The first attempt will be a non-final chunk. It should be answered
				// with a 308 to keep the upload retry process going, which allows the ChunkRetryDeadline timeout to fire.
				{byteRange: "bytes 0-14/*", responseStatus: 308, delay: 50 * time.Millisecond},
				{byteRange: "bytes 0-14/*", responseStatus: 308, delay: 50 * time.Millisecond},
				{byteRange: "bytes 0-14/*", responseStatus: 308, delay: 50 * time.Millisecond},
				{byteRange: "bytes 0-14/*", responseStatus: 308, delay: 50 * time.Millisecond},
			},
			shouldFail: true,
			wantError:  context.DeadlineExceeded,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// A media source that is slow to provide its data.
			media := &slowMediaReader{
				r:     strings.NewReader(data),
				delay: tc.mediaReadDelay,
			}

			// A transport that is slow to respond.
			transport := &interruptibleTransport{
				events: tc.events,
				bodies: bodyTracker{},
			}

			rx := &ResumableUpload{
				Client:               &http.Client{Transport: transport},
				Media:                NewMediaBuffer(media, len(data), false), // Chunk size is the whole payload.
				MediaType:            "text/plain",
				ChunkTransferTimeout: tc.chunkTransferTimeout,
				ChunkRetryDeadline:   100 * time.Millisecond,
			}

			// Use a backoff with no pause to speed up retries.
			oldBackoff := backoff
			backoff = func() Backoff { return new(NoPauseBackoff) }
			defer func() { backoff = oldBackoff }()

			res, err := rx.Upload(context.Background())
			if res != nil {
				res.Body.Close()
			}

			if tc.shouldFail {
				if err == nil {
					t.Fatalf("expected upload to fail, but it succeeded")
				}
				if !errors.Is(err, tc.wantError) {
					t.Fatalf("got error: %v, want error of type: %v", err, tc.wantError)
				}
			} else {
				if err != nil {
					t.Fatalf("expected upload to succeed, but it failed with: %v", err)
				}
				if res.StatusCode != http.StatusOK {
					t.Fatalf("got status code: %d, want: %d", res.StatusCode, http.StatusOK)
				}
			}
		})
	}
}

// failingReader returns an I/O error after a configurable number of bytes.
type failingReader struct {
	r         io.Reader
	failAfter int64
	read      int64
	err       error
}

func (r *failingReader) Read(p []byte) (n int, err error) {
	if r.read >= r.failAfter {
		return 0, r.err
	}
	n, err = r.r.Read(p)
	r.read += int64(n)
	if err != nil {
		return n, err
	}
	return n, nil
}

// TestMediaError verifies that an I/O error from the media source during an
// upload correctly terminates the process and propagates the error.
func TestMediaError(t *testing.T) {
	const (
		chunkSize = 100
		mediaSize = 300
		failAfter = int64(100)
	)
	baseReader := strings.NewReader(strings.Repeat("b", mediaSize))
	failErr := errors.New("simulated media I/O error")
	media := &failingReader{
		r:         baseReader,
		failAfter: failAfter,
		err:       failErr,
	}

	tr := &interruptibleTransport{
		events: []event{
			{byteRange: "bytes 0-99/*", responseStatus: 308},
		},
		bodies: bodyTracker{},
	}

	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewMediaBuffer(media, chunkSize, false),
		MediaType: "text/plain",
	}

	// Use a backoff with no pause to speed up retries.
	oldBackoff := backoff
	backoff = func() Backoff { return new(NoPauseBackoff) }
	defer func() { backoff = oldBackoff }()

	_, err := rx.Upload(context.Background())
	if !errors.Is(err, failErr) {
		t.Fatalf("Upload err: got %v; want %v", err, failErr)
	}

	// Verify that only the first chunk was transferred.
	if got, want := len(tr.buf), 100; got != want {
		t.Errorf("transferred bytes: got %d, want %d", got, want)
	}
}

// TestNonRetryableError verifies that the upload fails immediately without
// retrying when a non-retryable HTTP status code is received.
func TestNonRetryableError(t *testing.T) {
	const (
		chunkSize = 100
		mediaSize = 300
	)
	media := strings.NewReader(strings.Repeat("c", mediaSize))

	tr := &interruptibleTransport{
		events: []event{
			// The first request receives a 404, which is not retryable.
			{byteRange: "bytes 0-99/*", responseStatus: http.StatusNotFound},
		},
		bodies: bodyTracker{},
	}

	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewMediaBuffer(media, chunkSize, false),
		MediaType: "text/plain",
	}

	// Use a backoff with no pause to speed up retries.
	oldBackoff := backoff
	backoff = func() Backoff { return new(NoPauseBackoff) }
	defer func() { backoff = oldBackoff }()

	res, err := rx.Upload(context.Background())
	if res != nil {
		res.Body.Close()
	}

	if err == nil && res.StatusCode == http.StatusOK {
		t.Fatalf("expected upload to fail, but it succeeded")
	}
	// Check that no retries were attempted.
	if rx.attempts > 1 {
		t.Errorf("expected 1 attempt, but got %d", rx.attempts)
	}
	// Check that the transport has no leftover events, proving we stopped immediately.
	if len(tr.events) > 0 {
		t.Errorf("did not expect leftover events, but found %d", len(tr.events))
	}
	if len(tr.bodies) > 0 {
		t.Errorf("unclosed request bodies: %v", tr.bodies)
	}
}

func TestOverallUploadTimeout(t *testing.T) {
	const (
		chunkSize = 90
		mediaSize = 300
	)
	media := strings.NewReader(strings.Repeat("a", mediaSize))

	tr := &interruptibleTransport{
		buf: make([]byte, 0, mediaSize),
		events: []event{
			{byteRange: "bytes 0-89/*", responseStatus: 308},
			{byteRange: "bytes 90-179/*", responseStatus: 308, delay: 100 * time.Millisecond}, // This will cause a timeout
		},
		bodies: bodyTracker{},
	}

	rx := &ResumableUpload{
		Client:    &http.Client{Transport: tr},
		Media:     NewMediaBuffer(media, chunkSize, false),
		MediaType: "text/plain",
	}

	oldBackoff := backoff
	backoff = func() Backoff { return new(NoPauseBackoff) }
	defer func() { backoff = oldBackoff }()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	// The second event has a 100ms delay, so the upload is guaranteed to be
	// timed out by before it completes.
	_, err := rx.Upload(ctx)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Upload err: got: %v; want: context.DeadlineExceeded", err)
	}
}

func TestUploadChecksum(t *testing.T) {
	data := string(bytes.Repeat([]byte("a"), 300))
	chunkSize := 90
	checksum := crc32.Checksum([]byte(data), crc32cTable)
	tests := []struct {
		name               string
		chunkSize          int
		sendChecksum       bool
		wantChecksumHeader string
	}{
		{
			name:         "checksum disabled",
			sendChecksum: false,
		},
		{
			name:               "checksum enabled",
			sendChecksum:       true,
			wantChecksumHeader: fmt.Sprintf("%v=%v", crc32cPrefix, encodeUint32(checksum)),
		},
	}
	for _, tc := range tests {
		media := strings.NewReader(data)

		// Simulate multi-chunk resumable requests.
		tr := &interruptibleTransport{
			events: []event{
				{byteRange: "bytes 0-89/*", responseStatus: 308},
				{byteRange: "bytes 90-179/*", responseStatus: 308},
				{byteRange: "bytes 180-269/*", responseStatus: 308},
				{byteRange: "bytes 270-299/300", responseStatus: 200},
			},
			bodies: bodyTracker{},
		}
		rx := &ResumableUpload{
			Client:    &http.Client{Transport: tr},
			Media:     NewMediaBuffer(media, chunkSize, tc.sendChecksum),
			MediaType: "text/plain",
		}

		res, err := rx.Upload(context.Background())
		if err != nil {
			t.Fatalf("Upload failed: %v", err)
		}
		res.Body.Close()
		if gotChecksumHeader := tr.finalHeader.Get(hashHeaderKey); gotChecksumHeader != tc.wantChecksumHeader {
			t.Errorf("Hash header: got %q, want %q", gotChecksumHeader, tc.wantChecksumHeader)
		}
	}
}
