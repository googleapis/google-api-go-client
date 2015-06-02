// Copyright (C) 2015 Motorola Mobility LLC. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"errors"
	"io"
)

// sequentialChunker provides access to underlying content in a possibly overlapping
// but monotonically non-decreasing sequence of chunks, buffering content as necessary.
type sequentialChunker interface {
	// ChunkAt returns a io.Reader over a chunk of n bytes starting at given offset
	// within the underlying content. The returned error is:
	//  - nil if there is more content beyond this chunk,
	//  - io.EOF if this is the last chunk of the content,
	//  - an error value encountered while reading the underlying content.
	// ChunkAt is not required to support skipping content, nor backtracking prior to
	// the start of the previously read chunk. If one of those operations is attempted,
	// ChunkAt may return errSkip or errBackTrack, respectively.
	// When underlying reader size is known up-front, ChunkAt may return io.ErrUnexpectedEOF
	// if premature EOF is encountered.
	// The returned reader is valid only until the next ChunkAt call.
	// Advancing the offset via ChunkAt calls should release all buffered content prior to the
	// newly specified offset.
	ChunkAt(off int64) (r io.Reader, n int64, err error)
	// Size returns full content size if known, or 0 if size is unknown.
	Size() int64
	getChunkSize() int64 // private method for testing
}

var (
	errBacktrack = errors.New("sequentialChunker: advanced past the requested offset")
	errSkip      = errors.New("sequentialChunker: requested offset would cause skipping")
)

// sizedChunker is an implementation of sequentialChunker that wraps a io.ReaderAt
// with known content size. If chunkSize is 0, sizedChunker will always return a chunk
// corresponding to the entire remaining content.
type sizedChunker struct {
	r         io.ReaderAt
	size      int64 // content size
	chunkSize int64 // chunk size or 0 if chunking is off
}

func (sc *sizedChunker) ChunkAt(off int64) (io.Reader, int64, error) {
	n := sc.size - off
	if n <= 0 {
		return nil, 0, io.EOF
	}
	err := io.EOF
	if sc.chunkSize > 0 && sc.chunkSize < n {
		n = sc.chunkSize
		err = nil
	}
	return io.NewSectionReader(sc.r, off, n), n, err
}

func (sc *sizedChunker) Size() int64 { return sc.size }

func (sc *sizedChunker) getChunkSize() int64 { return sc.chunkSize }

// bufferedChunker is an implementation of sequentialChunker that buffers one chunk
// at the time from the underlying io.Reader.
type bufferedChunker struct {
	r         io.Reader // underlying reader
	chunkSize int64
	size      int64  // total content size, if known; 0 otherwise
	buf       []byte // buffered data from reader
	len       int    // number of bytes in buf
	pos       int64  // offset in underlying reader
}

func newBufferedChunker(r io.Reader, chunkSize, size int64) *bufferedChunker {
	if size > 0 {
		r = io.LimitReader(r, size)
	}
	return &bufferedChunker{r: r, chunkSize: chunkSize, size: size, buf: make([]byte, chunkSize+1)}
}

// ChunkAt returns a io.Reader over a chunk of n bytes starting at given offset
// within the underlying content. The returned error is:
//  - nil if there is more content beyond this chunk,
//  - io.EOF if this is the last chunk of the content,
//  - errSkip if specified offset would require skipping over unread content,
//  - errBackTrack if off is negative or lesser than offset of previously returned chunk,
//  - io.ErrUnexpectedEOF when premature EOF is encountered on a reader of known size, or
//  - an error value encountered while reading the underlying content.
// The returned reader is valid only until the next ChunkAt call.
// Advancing the offset via ChunkAt calls will release any buffered content prior to the
// newly specified offset.
func (bc *bufferedChunker) ChunkAt(off int64) (io.Reader, int64, error) {
	if off < bc.pos {
		return nil, 0, errBacktrack
	}
	discard := off - bc.pos
	if discard > int64(bc.len) {
		return nil, 0, errSkip
	}
	bc.len = copy(bc.buf, bc.buf[discard:bc.len])
	bc.pos = off
	// Extra bytes to accomodate, including 1 extra to detect EOF.
	if extra := int(bc.chunkSize + 1 - int64(bc.len)); extra > 0 {
		m, err := io.ReadFull(bc.r, bc.buf[bc.len:bc.len+extra])
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return nil, 0, err
		}
		bc.len += m
	}
	err := io.EOF
	n := int64(bc.len)
	if n > bc.chunkSize {
		n = bc.chunkSize
		err = nil
	}
	if bc.size > 0 && err == io.EOF && bc.pos+n != bc.size {
		return nil, 0, io.ErrUnexpectedEOF
	}
	return bytes.NewReader(bc.buf[:n]), n, err

}

func (bc *bufferedChunker) Size() int64 { return bc.size }

func (bc *bufferedChunker) getChunkSize() int64 { return bc.chunkSize }
