// Copyright (C) 2015 Motorola Mobility LLC All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"io"
	"io/ioutil"
	"runtime"
	"testing"
)

func setup100() *bufferedChunker {
	data := make([]byte, 100)
	for i := byte(0); i < 100; i++ {
		data[i] = i
	}
	r := bytes.NewReader(data)
	return newBufferedChunker(r, 100, 0)
}

func expect(t *testing.T, bc *bufferedChunker, buf []byte, off int64, n int64, err error, first, last byte, msg string) {
	bc.chunkSize = int64(len(buf))
	r, n1, err1 := bc.ChunkAt(off)
	if n != n1 {
		t.Fatalf("%s: expected %v bytes got %v, err %v", msg, n, n1, err1)
	}
	if err1 != err {
		t.Fatalf("%s: expected err %v got %v", msg, err, err1)
	}
	if n1 <= 0 {
		return
	}
	io.ReadFull(r, buf)
	if buf[0] != first || buf[len(buf)-1] != last {
		t.Fatalf("%s: wrong buf content, expected [%v...%v] got %v", msg, first, last, buf)
	}
}

func TestReadSkipRead(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 10)
	expect(t, bc, buf, 0, 10, nil, 0, 9, "initial read")
	expect(t, bc, buf, 20, 0, errSkip, 0, 9, "skip read 1")
	expect(t, bc, buf, 10, 10, nil, 10, 19, "read after skip")
}

func TestSkipInit(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 10)
	expect(t, bc, buf, 20, 0, errSkip, 0, 0, "skip read")
}

func TestBacktrack(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 10)
	expect(t, bc, buf, 0, 10, nil, 0, 9, "initial read")
	expect(t, bc, buf, 10, 10, nil, 10, 19, "second read")
	expect(t, bc, buf[0:5], 13, 5, nil, 13, 17, "back tracked read")
	expect(t, bc, buf, 15, 10, nil, 15, 24, "back tracked + new read")
	expect(t, bc, buf[0:5], 20, 5, nil, 20, 24, "back tracked to the end read")
}

func TestBacktrackErr(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 10)
	expect(t, bc, buf, 0, 10, nil, 0, 9, "initial read")
	expect(t, bc, buf, 10, 10, nil, 10, 19, "second read")
	expect(t, bc, buf, 9, 0, errBacktrack, buf[0], buf[9], "backtrack error read")
}

func TestEOFErr(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 50)
	expect(t, bc, buf, 0, 50, nil, 0, 49, "initial read")
	expect(t, bc, buf, 50, 50, io.EOF, 50, 99, "second read")
}

func TestEOFErr2(t *testing.T) {
	bc := setup100()
	buf := make([]byte, 51)
	expect(t, bc, buf, 0, 51, nil, 0, 50, "initial read")
	expect(t, bc, buf, 51, 49, io.EOF, 51, buf[50], "second read")
}

type endlessReader int

func (r endlessReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = byte(i)
	}
	return len(p), nil
}

// TestMemProfile emulates the way resumable upload would read sequence of chunks,
// while checking that excess memory allocations don't exceed twice the chunk size.
func TestMemProfile(t *testing.T) {
	const (
		chunkSize = 10 << 20
		numChunks = 20
	)

	allocated := func() uint64 { // current allocations after GC
		runtime.GC()
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		return m.Alloc
	}

	alloc := allocated()
	c := newBufferedChunker(endlessReader(0), chunkSize, numChunks*chunkSize)
	for start := int64(0); ; {
		r, size, err := c.ChunkAt(start)
		if size != chunkSize {
			t.Fatalf("size = %v, want %v", size, chunkSize)
		}
		start += size
		n, err1 := io.CopyN(ioutil.Discard, r, size)
		if err != io.EOF && err != nil {
			t.Fatalf("err = %v; want %v or %v", err1, nil, io.EOF)
		}
		if n != size {
			t.Fatalf("n = %v, want %v", n, size)
		}
		if alloc1 := allocated(); alloc1 > alloc+2*chunkSize {
			t.Fatalf("Allocations increased to more than 2 chunkSize: %v", alloc1-alloc)
		}
		if err == io.EOF {
			break
		}
	}
}
