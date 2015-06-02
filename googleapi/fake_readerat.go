// Copyright (C) 2015 Motorola Mobility LLC All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"errors"
	"io"
	"io/ioutil"
)

type fakeReaderAt struct {
	r   io.Reader // underlying reader
	buf []byte    // current buffer, len(buf) is how much data we hold
	pos int64     // position in underlying reader where current buffer starts
}

// FakeReaderAt adheres to io.ReaderAt interface, but with a different contract.
// Normally, ReadAt(slice, offset) can be used to perform random reads (offset can
// "jump" back and forth in an arbitrary way). On the other hand, FakeReaderAt
// implements io.ReaderAt via limited buffering of the underlying stream io.Reader,
// such that in a sequence of two ReadAt calls:
//
//    fr.ReadAt(s1, offset1)
//    fr.ReadAt(s2, offset2)
//
// the two offsets must satisfy: offset2 >= offset1.
//
// The interface also provides SizeAt call which should do reading from and buffering
// the underlying stream, and return the same result that ReadAt() would have
// returned if called, but without requiring caller to provide a []byte to copy
// the data into. Furthermore, the EOF semantics is stricter than for the normal
// io.ReaderAt: FakeReaderAt is guaranteed to return io.EOF along with the last byte
// of content from the underlying stream. In other words, if ReadAt() or SizeAt() return
// err = nil, that means that there is at least one more byte to be read.
//
// This odd set of behaviors is motivated by the needs of the resumable upload
// implementation, which expects io.ReaderAt, but really only requires buffering
// single chunk at a time, and which also needs to know that it's sending the last
// chunk at the time when it's sending it, as opposed to discovering EOF at the time
// of the next read.
type FakeReaderAt interface {
	io.ReaderAt
	SizeAt(off int64, size int) (n int, err error)
}

var (
	ErrBacktrack = errors.New("FakeReaderAt advanced past the requested offset")
)

func NewFakeReaderAt(r io.Reader) FakeReaderAt {
	return &fakeReaderAt{r: r}
}

func (r *fakeReaderAt) SizeAt(off int64, size int) (n int, err error) {
	if off < r.pos {
		return 0, ErrBacktrack
	}
	if off < r.pos+int64(len(r.buf)) { // offset falls within the buffered content
		// advance the buffer to offset position
		r.buf = r.buf[off-r.pos:]
		r.pos = off
		if size+1 <= len(r.buf) {
			return size, nil
		}
		newBuf := make([]byte, len(r.buf), size+1) // ask for 1 extra byte to detect EOF
		copy(newBuf, r.buf)
		r.buf = newBuf
		toRead := cap(r.buf) - len(r.buf)
		n1, err1 := io.ReadFull(r.r, r.buf[len(r.buf):cap(r.buf)])
		r.buf = r.buf[:len(r.buf)+n1]
		if n1 == toRead {
			return len(r.buf) - 1, nil // don't report the extra byte in the length
		}
		if err1 == io.ErrUnexpectedEOF {
			err1 = io.EOF
		}
		return len(r.buf), err1
	}
	if off > r.pos+int64(len(r.buf)) {
		// skip until offset
		_, err1 := io.CopyN(ioutil.Discard, r.r, off-(r.pos+int64(len(r.buf))))
		if err1 != nil {
			return 0, err1
		}
	}
	r.buf = make([]byte, size+1) // ask for 1 extra byte to detect EOF
	n1, err1 := io.ReadFull(r.r, r.buf)
	r.buf = r.buf[:n1]
	r.pos = off
	if n1 == size+1 {
		return size, nil
	}
	if err1 == io.ErrUnexpectedEOF {
		err1 = io.EOF
	}
	return n1, err1
}

func (r *fakeReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n1, err1 := r.SizeAt(off, len(p))
	if n1 > 0 {
		copy(p[0:n1], r.buf)
	}
	return n1, err1
}
