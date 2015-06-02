// Copyright 2015 The Go Authors
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"io"
	"testing"
)

// buffered100 returns new bufferedChunker with underlying content being a
// sequence of bytes: 0..99.
func buffered100(chunkSize int64) *bufferedChunker {
	data := make([]byte, 100)
	for i := 0; i < 100; i++ {
		data[i] = byte(i)
	}
	r := bytes.NewReader(data)
	return &bufferedChunker{r: r, chunkSize: chunkSize}
}

// sized100 returns new sizedChunker with underlying content being a
// sequence of bytes: 0..99.
func sized100(chunkSize int64) *sizedChunker {
	data := make([]byte, 100)
	for i := 0; i < 100; i++ {
		data[i] = byte(i)
	}
	r := bytes.NewReader(data)
	return &sizedChunker{r: r, chunkSize: chunkSize, size: 100}
}

// check verifies that ChunkAt returns expected values, returning the reader over the chunk.
func check(t *testing.T, sc sequentialChunker, off int64, n int64, err error) io.Reader {
	r, n1, err1 := sc.ChunkAt(off)
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
// returned reader, checking first and last byte of content and its size.
func checkContent(t *testing.T, sc sequentialChunker, off int64, n int64, err error, first byte, last byte) {
	r := check(t, sc, off, n, err)
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
	bc1 := buffered100(10)
	bc2 := buffered100(10)
	t.Log("Initial read")
	checkContent(t, bc1, 0, 10, nil, 0, 9)
	checkContent(t, bc2, 0, 10, nil, 0, 9)
	t.Log("Valid next read")
	checkContent(t, bc1, 10, 10, nil, 10, 19)
	t.Log("Skip read")
	check(t, bc2, 11, 0, errSkip)
}

func TestSkipInit(t *testing.T) {
	bc := buffered100(10)
	check(t, bc, 20, 0, errSkip)
}

func TestOverlap(t *testing.T) {
	bc := buffered100(10)
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
	bc := buffered100(10)
	t.Log("First read")
	checkContent(t, bc, 0, 10, nil, 0, 9)
	t.Log("Second read")
	checkContent(t, bc, 10, 10, nil, 10, 19)
	t.Log("Backtracked read")
	check(t, bc, 9, 0, errBacktrack)
}

func TestBacktrackErrInit(t *testing.T) {
	bc := buffered100(10)
	check(t, bc, -1, 0, errBacktrack)
}

func TestEOFErr(t *testing.T) {
	bc := buffered100(50)
	t.Log("Initial read")
	checkContent(t, bc, 0, 50, nil, 0, 49)
	t.Log("Final read")
	checkContent(t, bc, 50, 50, io.EOF, 50, 99)
}

func TestShortChunk(t *testing.T) {
	bc := buffered100(99) // chunkSize = contentSize - 1
	t.Log("Initial read")
	checkContent(t, bc, 0, 99, nil, 0, 98)
	t.Log("Final read of one byte")
	checkContent(t, bc, 99, 1, io.EOF, 99, 99)
}

func TestEqualChunk(t *testing.T) {
	bc := buffered100(100) // chunkSize = contentSize
	checkContent(t, bc, 0, 100, io.EOF, 0, 99)
}

func TestLongChunk(t *testing.T) {
	bc := buffered100(101) // chunkSize = contentSize + 1
	checkContent(t, bc, 0, 100, io.EOF, 0, 99)
}

func TestEmptyContent(t *testing.T) {
	var empty []byte
	bc := &bufferedChunker{r: bytes.NewReader(empty), chunkSize: 10}
	check(t, bc, 0, 0, io.EOF)
	check(t, bc, 0, 0, io.EOF) // second time should get the same result
}

func TestSizedChunker(t *testing.T) {
	sc := sized100(0) // no chunking
	t.Log("Initial read")
	checkContent(t, sc, 0, 100, io.EOF, 0, 99)
	t.Log("Complete overlap")
	checkContent(t, sc, 0, 100, io.EOF, 0, 99)
	t.Log("Partial overlap")
	checkContent(t, sc, 1, 99, io.EOF, 1, 99)
}

func TestSizedChunkerShortChunk(t *testing.T) {
	sc := sized100(99) // chunkSize = contentSize - 1
	checkContent(t, sc, 0, 99, nil, 0, 98)
	checkContent(t, sc, 99, 1, io.EOF, 99, 99)
}

func TestSizedChunkerEqualChunk(t *testing.T) {
	sc := sized100(100) // chunkSize = contentSize
	checkContent(t, sc, 0, 100, io.EOF, 0, 99)
}

func TestSizedChunkerLongChunk(t *testing.T) {
	sc := sized100(101) // chunkSize = contentSize + 1
	checkContent(t, sc, 0, 100, io.EOF, 0, 99)
}

func TestSizedChunkerEmptyContent(t *testing.T) {
	var empty []byte
	sc := &sizedChunker{r: bytes.NewReader(empty), size: 0}
	check(t, sc, 0, 0, io.EOF)
	check(t, sc, 0, 0, io.EOF) // second time should get the same result
}
