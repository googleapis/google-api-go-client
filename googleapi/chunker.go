// Copyright (C) 2015 Motorola Mobility LLC. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
)

// sequentialChunker provides access to underlying content in a possibly overlapping
// but monotonically non-decreasing sequence of chunks, buffering content as necessary.
type sequentialChunker interface {
	// ChunkAt returns a io.Reader over a chunk of n bytes starting at given offset
	// within the underlying content. The returned error is:
	//  - nil if there is more content beyond this chunk,
	//  - io.EOF if this is the last chunk of the content,
	//  - an error value encountered while reading the underlying content.
	// Chunker is not required to support skipping content, nor backtracking prior to
	// the start of the previously read chunk. If one of those operations is attempted,
	// chunker may return errSkip or errBackTrack, respectively.
	// The returned reader is valid only until the next ChunkAt call.
	ChunkAt(off int64) (r io.Reader, n int64, err error)
	// Size returns full content size if known, or 0 if size is unknown.
	Size() int64
	getChunkSize() int64 // private method for testing
}

var (
	errBacktrack = errors.New("sequentialChunker advanced past the requested offset")
	errSkip      = errors.New("sequentialChunker: requested offset would cause skipping")
)

// sizedChunker is an implementation of sequentialChunker that wraps a io.ReaderAt
// with known content size. If chunkSize is 0, sizedChunker will always return chunk
// corresponding to the entire remaining content.
type sizedChunker struct {
	r         io.ReaderAt
	size      int64 // content size
	chunkSize int64 // chunk size or 0 if chunking is off
}

func (sc *sizedChunker) ChunkAt(off int64) (r io.Reader, n int64, err error) {
	err = io.EOF
	if n = sc.size - off; n <= 0 {
		n = 0
		return
	}
	if sc.chunkSize > 0 && sc.chunkSize < n {
		n = sc.chunkSize
		err = nil
	}
	r = io.NewSectionReader(sc.r, off, n)
	return
}

func (sc *sizedChunker) Size() int64 { return sc.size }

func (sc *sizedChunker) getChunkSize() int64 { return sc.chunkSize }

// bufferedChunker is an implementation of sequentialChunker that buffers the content
// from the underlying io.Reader. Advancing the offset via ChunkAt calls causes
// previously buffered content to be released.
type bufferedChunker struct {
	r         io.Reader // underlying reader
	chunkSize int64
	size      int64        // total content size, if known; 0 otherwise
	buf       bytes.Buffer // buffered data from reader
	pos       int64        // offset in underlying reader
	pos2      int64        // offset of first unread byte
}

func newBufferedChunker(r io.Reader, chunkSize, size int64) *bufferedChunker {
	if size > 0 {
		r = io.LimitReader(r, size)
	}
	return &bufferedChunker{r: r, chunkSize: chunkSize, size: size}
}

func (bc *bufferedChunker) ChunkAt(off int64) (r io.Reader, n int64, err error) {
	if off < bc.pos {
		err = errBacktrack
		return
	}
	if off > bc.pos2 {
		err = errSkip
		return
	}
	// Discard *prior* to buffering new content, to facilitate memory release.
	if discard := off - bc.pos; discard > 0 {
		io.CopyN(ioutil.Discard, &bc.buf, discard)
		bc.pos = off
	}
	// Extra bytes to accomodate, including 1 extra to detect EOF.
	if extra := bc.chunkSize + 1 - int64(bc.buf.Len()); extra > 0 {
		_, err = io.CopyN(&bc.buf, bc.r, extra)
	}
	r = bytes.NewReader(bc.buf.Bytes())
	if n = int64(bc.buf.Len()); n > bc.chunkSize {
		n = bc.chunkSize
		r = io.LimitReader(r, n)
	}
	bc.pos2 = bc.pos + n
	return
}

func (bc *bufferedChunker) Size() int64 { return bc.size }

func (bc *bufferedChunker) getChunkSize() int64 { return bc.chunkSize }
