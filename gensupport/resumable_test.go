// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type unexpectedReader struct{}

func (unexpectedReader) Read([]byte) (int, error) {
	return 0, fmt.Errorf("unexpected read in test")
}

var contentRangeRE = regexp.MustCompile(`^bytes (\d+)\-(\d+)/(\d+)$`)

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req
	if rng := req.Header.Get("Content-Range"); rng != "" && !strings.HasPrefix(rng, "bytes */") { // Read the data
		m := contentRangeRE.FindStringSubmatch(rng)
		if len(m) != 4 {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		start, err := strconv.ParseInt(m[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		end, err := strconv.ParseInt(m[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		totalSize, err := strconv.ParseInt(m[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		partialSize := end - start + 1
		t.buf, err = ioutil.ReadAll(req.Body)
		if err != nil || int64(len(t.buf)) != partialSize {
			return nil, fmt.Errorf("unable to read %v bytes from request data, n=%v: %v", partialSize, len(t.buf), err)
		}
		if totalSize == end+1 {
			t.statusCode = 200 // signify completion of transfer
		}
	}
	f := ioutil.NopCloser(unexpectedReader{})
	res := &http.Response{
		Body:       f,
		StatusCode: t.statusCode,
		Header:     http.Header{},
	}
	if t.rangeVal != "" {
		res.Header.Set("Range", t.rangeVal)
	}
	return res, nil
}

type testTransport struct {
	req        *http.Request
	statusCode int
	rangeVal   string
	want       int64
	buf        []byte
}

var statusTests = []*testTransport{
	&testTransport{statusCode: 308, want: 0},
	&testTransport{statusCode: 308, rangeVal: "bytes=0-0", want: 1},
	&testTransport{statusCode: 308, rangeVal: "bytes=0-42", want: 43},
}

func TestTransferStatus(t *testing.T) {
	ctx := context.Background()
	for _, tr := range statusTests {
		rx := &ResumableUpload{
			Client: &http.Client{Transport: tr},
		}
		g, _, err := rx.transferStatus(ctx)
		if err != nil {
			t.Error(err)
		}
		if g != tr.want {
			t.Errorf("transferStatus got %v, want %v", g, tr.want)
		}
	}
}

func (t *interruptedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req
	if rng := req.Header.Get("Content-Range"); rng != "" && !strings.HasPrefix(rng, "bytes */") {
		t.interruptCount += 1
		if t.interruptCount%7 == 0 { // Respond with a "service unavailable" error
			res := &http.Response{
				StatusCode: http.StatusServiceUnavailable,
				Header:     http.Header{},
			}
			t.rangeVal = fmt.Sprintf("bytes=0-%v", len(t.buf)-1) // Set the response for next time
			return res, nil
		}
		m := contentRangeRE.FindStringSubmatch(rng)
		if len(m) != 4 {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		start, err := strconv.ParseInt(m[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		end, err := strconv.ParseInt(m[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		totalSize, err := strconv.ParseInt(m[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse content range: %v", rng)
		}
		partialSize := end - start + 1
		buf, err := ioutil.ReadAll(req.Body)
		if err != nil || int64(len(buf)) != partialSize {
			return nil, fmt.Errorf("unable to read %v bytes from request data, n=%v: %v", partialSize, len(buf), err)
		}
		t.buf = append(t.buf, buf...)
		if totalSize == end+1 {
			t.statusCode = 200 // signify completion of transfer
		}
	}
	f := ioutil.NopCloser(unexpectedReader{})
	res := &http.Response{
		Body:       f,
		StatusCode: t.statusCode,
		Header:     http.Header{},
	}
	if t.rangeVal != "" {
		res.Header.Set("Range", t.rangeVal)
	}
	return res, nil
}

type interruptedTransport struct {
	req             *http.Request
	statusCode      int
	rangeVal        string
	interruptCount  int
	buf             []byte
	progressUpdates []int64
}

func (tr *interruptedTransport) ProgressUpdate(current int64) {
	tr.progressUpdates = append(tr.progressUpdates, current)
}

func TestInterruptedTransferChunks(t *testing.T) {
	// TODO(mcgreevy): don't read from the filesystem here.
	f, err := os.Open("resumable.go")
	if err != nil {
		t.Fatalf("unable to open resumable.go: %v", err)
	}
	defer f.Close()
	slurp, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("unable to slurp file: %v", err)
	}
	st, err := f.Stat()
	if err != nil {
		t.Fatalf("unable to stat resumable.go: %v", err)
	}
	tr := &interruptedTransport{
		statusCode: 308,
		buf:        make([]byte, 0, st.Size()),
	}
	oldChunkSize := chunkSize
	defer func() { chunkSize = oldChunkSize }()
	chunkSize = 100 // override to process small chunks for test.

	sleep = func(time.Duration) {} // override time.Sleep
	rx := &ResumableUpload{
		Client:        &http.Client{Transport: tr},
		Media:         f,
		MediaType:     "text/plain",
		ContentLength: st.Size(),
		Callback:      tr.ProgressUpdate,
	}
	res, err := rx.Upload(context.Background())
	if err != nil || res == nil || res.StatusCode != http.StatusOK {
		if res == nil {
			t.Errorf("transferChunks not successful, res=nil: %v", err)
		} else {
			t.Errorf("transferChunks not successful, statusCode=%v: %v", res.StatusCode, err)
		}
	}
	if len(tr.buf) != len(slurp) || bytes.Compare(tr.buf, slurp) != 0 {
		t.Errorf("transferred file corrupted:\ngot %s\nwant %s", tr.buf, slurp)
	}
	want := []int64{}
	for i := chunkSize; i <= st.Size(); i += chunkSize {
		want = append(want, i)
	}
	if st.Size()%chunkSize != 0 {
		want = append(want, st.Size())
	}
	if !reflect.DeepEqual(tr.progressUpdates, want) {
		t.Errorf("progress update error, got %v, want %v", tr.progressUpdates, want)
	}
}

func TestCancelUpload(t *testing.T) {
	f, err := os.Open("resumable.go")
	if err != nil {
		t.Fatalf("unable to open resumable.go: %v", err)
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		t.Fatalf("unable to stat resumable.go: %v", err)
	}
	tr := &interruptedTransport{
		statusCode: 308,
		buf:        make([]byte, 0, st.Size()),
	}
	oldChunkSize := chunkSize
	defer func() { chunkSize = oldChunkSize }()
	chunkSize = 100 // override to process small chunks for test.

	sleep = func(time.Duration) {} // override time.Sleep
	rx := &ResumableUpload{
		Client:        &http.Client{Transport: tr},
		Media:         f,
		MediaType:     "text/plain",
		ContentLength: st.Size(),
		Callback:      tr.ProgressUpdate,
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc() // stop the upload that hasn't started yet
	res, err := rx.Upload(ctx)
	if err == nil || res == nil || res.StatusCode != http.StatusRequestTimeout {
		if res == nil {
			t.Errorf("transferChunks not successful, got res=nil, err=%v, want StatusRequestTimeout", err)
		} else {
			t.Errorf("transferChunks not successful, got statusCode=%v, err=%v, want StatusRequestTimeout", res.StatusCode, err)
		}
	}
}
