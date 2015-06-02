// Copyright 2015 The Go Authors
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
	// The returned reader is valid only until the next ChunkAt call, and may be nil when
	// non-nil error other than io.EOF is returned.
	// Advancing the offset via ChunkAt calls should release all buffered content prior to the
	// newly specified offset.
	ChunkAt(off int64) (r io.Reader, n int64, err error)
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
	if n < 0 { // an attempt to read past the end of content
		n = 0
		off = sc.size
	}
	err := io.EOF
	if sc.chunkSize > 0 && sc.chunkSize < n {
		n = sc.chunkSize
		err = nil
	}
	return io.NewSectionReader(sc.r, off, n), n, err
}

func (sc *sizedChunker) getChunkSize() int64 { return sc.chunkSize }

// bufferedChunker is an implementation of sequentialChunker that buffers one chunk
// at a time from the underlying io.Reader.
type bufferedChunker struct {
	r         io.Reader // underlying reader
	chunkSize int64
	buf       []byte // buffered data from reader
	len       int    // number of bytes in buf
	pos       int64  // offset in underlying reader
}

// ChunkAt returns a io.Reader over a chunk of n bytes starting at given offset
// within the underlying content. The returned error is:
//  - nil if there is more content beyond this chunk,
//  - io.EOF if this is the last chunk of the content,
//  - errSkip if specified offset would require skipping over unread content,
//  - errBackTrack if off is negative or lesser than offset of previously returned chunk, or
//  - an error value encountered while reading the underlying content.
// The returned reader is valid only until the next ChunkAt call.
// Advancing the offset via ChunkAt calls will release any buffered content prior to the
// newly specified offset.
func (bc *bufferedChunker) ChunkAt(off int64) (io.Reader, int64, error) {
	if bc.buf == nil {
		bc.buf = make([]byte, bc.chunkSize+1) // extra byte is for EOF detection
	}
	if off < bc.pos {
		return nil, 0, errBacktrack
	}
	discard := off - bc.pos
	if discard > int64(bc.len) || discard > bc.chunkSize {
		return nil, 0, errSkip
	}
	bc.len = copy(bc.buf, bc.buf[discard:bc.len])
	bc.pos = off
	// Refill the buffer.
	m, err := io.ReadFull(bc.r, bc.buf[bc.len:])
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, 0, err
	}
	bc.len += m
	err = io.EOF
	n := int64(bc.len)
	if n > bc.chunkSize { // really just bc.chunkSize+1
		n = bc.chunkSize
		err = nil
	}
	return bytes.NewReader(bc.buf[:n]), n, err

}

func (bc *bufferedChunker) getChunkSize() int64 { return bc.chunkSize }
