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

// setup100 returns new bufferedChunker with underlying content being a
// sequence of bytes: 0..99
func setup100(chunkSize, contentSize int64) *bufferedChunker {
	data := make([]byte, 100)
	for i := byte(0); i < 100; i++ {
		data[i] = i
	}
	r := bytes.NewReader(data)
	return newBufferedChunker(r, chunkSize, contentSize)
}

// check verifies that ChunkAt returns expected values, returning the reader over the chunk
func check(t *testing.T, bc *bufferedChunker, off int64, n int64, err error) io.Reader {
	r, n1, err1 := bc.ChunkAt(off)
	if n != n1 {
		t.Fatalf("size = %v, want %v", n1, n)
	}
	if err1 != err {
		t.Fatalf("err = %v, want %v", err1, err)
	}
	if v, w := (r == nil), (err1 != nil && err1 != io.EOF); v != w {
		t.Fatalf("isReaderNil = %v, want %v", v, w)
	}
	return r
}

// checkContent first peforms check() and then reads full content from the
// returned reader, checking first and last byte of content and its size
func checkContent(t *testing.T, bc *bufferedChunker, off int64, n int64, err error, first byte, last byte) {
	r := check(t, bc, off, n, err)
	if r == nil {
		t.Fatalf("reader = nil, want non-nil")
	}
	//make the buffer larger by 1 byte, so we detect if reader is too large
	buf := make([]byte, n+1)
	n1, _ := io.ReadFull(r, buf)
	if int64(n1) != n {
		t.Fatalf("reader size = %v, want %v", n1, n)
	}
	if buf[0] != first || buf[n-1] != last {
		t.Fatalf("wrong buf content, got %v, want [%v...%v]", buf[:n], first, last)
	}
}

func TestReadSkip(t *testing.T) {
	bc := setup100(10, 0)
	t.Log("Initial read")
	checkContent(t, bc, 0, 10, nil, 0, 9)
	t.Log("Skip read")
	check(t, bc, 20, 0, errSkip)
}

func TestSkipInit(t *testing.T) {
	bc := setup100(10, 0)
	check(t, bc, 20, 0, errSkip)
}

func TestOverlap(t *testing.T) {
	bc := setup100(10, 0)
	t.Log("Initial read")
	checkContent(t, bc, 0, 10, nil, 0, 9)
	t.Log("Second read")
	checkContent(t, bc, 10, 10, nil, 10, 19)
	t.Log("Exact overlap read")
	checkContent(t, bc, 10, 10, nil, 10, 19)
	t.Log("Partial overlap read")
	checkContent(t, bc, 13, 10, nil, 13, 22)
}

func TestBacktrackErr(t *testing.T) {
	bc := setup100(10, 0)
	t.Log("First read")
	checkContent(t, bc, 0, 10, nil, 0, 9)
	t.Log("Second read")
	checkContent(t, bc, 10, 10, nil, 10, 19)
	t.Log("Backtracked read")
	check(t, bc, 9, 0, errBacktrack)
}

func TestBacktrackErrInit(t *testing.T) {
	bc := setup100(10, 0)
	check(t, bc, -1, 0, errBacktrack)
}

func TestEOFErr(t *testing.T) {
	//
	bc := setup100(50, 0)
	t.Log("Initial read")
	checkContent(t, bc, 0, 50, nil, 0, 49)
	t.Log("Final read")
	checkContent(t, bc, 50, 50, io.EOF, 50, 99)
}

func TestShortChunk(t *testing.T) {
	bc := setup100(99, 0) // chunkSize = contentSize - 1
	t.Log("Initial read")
	checkContent(t, bc, 0, 99, nil, 0, 98)
	t.Log("Final read of one byte")
	checkContent(t, bc, 99, 1, io.EOF, 99, 99)
}

func TestEqualChunk(t *testing.T) {
	bc := setup100(100, 0) // chunkSize = contentSize
	checkContent(t, bc, 0, 100, io.EOF, 0, 99)
}

func TestLongChunk(t *testing.T) {
	bc := setup100(101, 0) // chunkSize = contentSize + 1
	checkContent(t, bc, 0, 100, io.EOF, 0, 99)
}

func TestShortContent(t *testing.T) {
	bc := setup100(200, 90)
	checkContent(t, bc, 0, 90, io.EOF, 0, 89)
}

func TestPrematureEOF(t *testing.T) {
	bc := setup100(200, 110)
	check(t, bc, 0, 0, io.ErrUnexpectedEOF)
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
