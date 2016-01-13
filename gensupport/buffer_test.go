// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"io"
	"io/ioutil"
	"reflect"
	"testing"
	"testing/iotest"
)

// getChunkAsString reads a chunk from rb, but does not call Next.
func getChunkAsString(t *testing.T, rb *ResumableBuffer) (string, error) {
	chunk, size, err := rb.Chunk()

	buf, e := ioutil.ReadAll(chunk)
	if e != nil {
		t.Fatalf("Failed reading chunk: %v", e)
	}
	if size != len(buf) {
		t.Fatalf("reported chunk size doesn't match actual chunk size: got: %v; want: %v", size, len(buf))
	}
	return string(buf), err
}

func TestChunking(t *testing.T) {
	type testCase struct {
		data       string // the data to read from the Reader
		finalErr   error  // error to return after data has been read
		chunkSize  int
		wantChunks []string
	}

	for _, singleByteReads := range []bool{true, false} {
		for _, tc := range []testCase{
			{
				data:       "abcdefg",
				finalErr:   nil,
				chunkSize:  3,
				wantChunks: []string{"abc", "def", "g"},
			},
			{
				data:       "abcdefg",
				finalErr:   nil,
				chunkSize:  1,
				wantChunks: []string{"a", "b", "c", "d", "e", "f", "g"},
			},
			{
				data:       "abcdefg",
				finalErr:   nil,
				chunkSize:  7,
				wantChunks: []string{"abcdefg"},
			},
			{
				data:       "abcdefg",
				finalErr:   nil,
				chunkSize:  8,
				wantChunks: []string{"abcdefg"},
			},
			{
				data:       "abcdefg",
				finalErr:   errBang,
				chunkSize:  3,
				wantChunks: []string{"abc", "def", "g"},
			},
			{
				data:       "abcdefg",
				finalErr:   errBang,
				chunkSize:  8,
				wantChunks: []string{"abcdefg"},
			},
		} {
			var r io.Reader = &errReader{buf: []byte(tc.data), err: tc.finalErr}

			if singleByteReads {
				r = iotest.OneByteReader(r)
			}

			rb := NewResumableBuffer(r, tc.chunkSize)
			var gotErr error
			got := []string{}
			for {
				chunk, err := getChunkAsString(t, rb)
				if len(chunk) != 0 {

					got = append(got, string(chunk))
				}
				if err != nil {
					gotErr = err
					break
				}
				rb.Next()
			}

			if !reflect.DeepEqual(got, tc.wantChunks) {
				t.Fatalf("Failed reading buffer: got: %v; want:%v", got, tc.wantChunks)
			}

			expectedErr := tc.finalErr
			if expectedErr == nil {
				expectedErr = io.EOF
			}
			if gotErr != expectedErr {
				t.Fatalf("Reading buffer error: got: %v; want: %v", gotErr, expectedErr)
			}
		}
	}
}

func TestChunkCanBeReused(t *testing.T) {
	er := &errReader{buf: []byte("abcdefg")}
	rb := NewResumableBuffer(er, 3)

	// expectChunk reads a chunk and checks that it got what was wanted.
	expectChunk := func(want string, wantErr error) {
		got, err := getChunkAsString(t, rb)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Failed reading buffer: got: %q; want:%q", got, want)
		}
		if err != wantErr {
			t.Fatalf("error reading buffer: got: %v; want: %v", err, wantErr)
		}
	}
	expectChunk("abc", nil)
	// On second call, should get same chunk again.
	expectChunk("abc", nil)
	rb.Next()
	expectChunk("def", nil)
	expectChunk("def", nil)
	rb.Next()
	expectChunk("g", io.EOF)
	expectChunk("g", io.EOF)
	rb.Next()
	expectChunk("", io.EOF)
}

func TestPos(t *testing.T) {
	er := &errReader{buf: []byte("abcdefg")}
	rb := NewResumableBuffer(er, 3)

	expectPos := func(want int64) {
		if got := rb.Pos(); got != want {
			t.Errorf("resumable buffer Pos: got: %v; want: %v", got, want)
		}
	}

	expectPos(0)
	// Simply reading in a chunk will not advance pos.
	rb.Chunk()
	expectPos(0)

	// Calling Next after Chunk will advance pos.
	rb.Next()
	expectPos(3)

	// Calling Next again will do nothing, because there is no current chunk.
	rb.Next()
	expectPos(3)

	rb.Chunk()
	expectPos(3)
	rb.Next()
	expectPos(6)
	rb.Next()
	expectPos(6)

	// Load the final 1-byte chunk.
	rb.Chunk()
	expectPos(6)
	// Next will advance 1 byte.  But there are no more chunks, so pos will not increase beyond 7.
	rb.Next()
	expectPos(7)
	rb.Chunk()
	expectPos(7)
	rb.Next()
	expectPos(7)
}
