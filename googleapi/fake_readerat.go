// Copyright (C) 2015 Motorola Mobility LLC. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"errors"
	"io"
)

// sequentialReaderAt embellishes io.ReaderAt with additional methods needed for
// resumable upload.
type sequentialReaderAt interface {
	io.ReaderAt
	// SizeAt returns number of bytes of content in the underlying reader, starting
	// at offset off and up to the provided size. EOF semantics is stricter than for
	// ReadAt: io.EOF is guaranteed to be returned *with* (as opposed to after)
	// the last byte of content.
	// Note that SizeAt may cause the relevant bytes to be read and buffered in
	// memory.
	// As a special case, passing size=0 will cause SizeAt to return full
	// content size starting with given offset and err=io.EOF provided that the
	// content size is known; otherwise err=errUnknownSize is returned.
	SizeAt(off, size int64) (n int64, err error)
	// ReleaseAt signals to the underlying reader that it is no longer necessary
	// to buffer content prior to the given offset. After ReleaseAt is called,
	// calls to ReadAt with lesser offset values may fail with errBackTrack.
	ReleaseAt(off int64) error
}

// sizedReaderAt is an implementation of sequentialReaderAt that wraps file-like
// reader: a io.ReaderAt with known content size. ReleaseAt is a no-op.
type sizedReaderAt struct {
	io.ReaderAt
	size int64
}

func (sr *sizedReaderAt) SizeAt(off, size int64) (n int64, err error) {
	if off < 0 || off >= sr.size {
		return 0, io.EOF
	}
	if size == 0 {
		// special case: size = 0 means report full remaining content size
		return sr.size - off, io.EOF
	}
	if sr.size > off+size {
		return size, nil
	}
	return sr.size - off, io.EOF
}

func (sr *sizedReaderAt) ReleaseAt(off int64) error {
	if off < 0 || off > sr.size {
		return errBadOffset
	}
	return nil
}

// fakeReaderAt is an implementation of sequentialReaderAt that buffers the content
// from the underlying io.Reader in order to provide io.ReaderAt interface.
// ReleaseAt is used to control the memory impact of buffering. When reading content,
// one extra byte is always requested from the underlying reader to ensure the strict
// EOF semantics of SizeAt method.
//
// fakeReaderAt was motivated by the needs of the resumable upload implementation,
// which expects io.ReaderAt, but really only requires buffering single chunk at a
// time, and which also needs to know that it's sending the last chunk at the time
// when it's sending it, as opposed to discovering EOF at the time of the next read.
type fakeReaderAt struct {
	r   io.Reader // underlying reader
	buf []byte    // current buffer, len(buf) is how much data we hold
	pos int64     // position in underlying reader where current buffer starts
}

var (
	errBacktrack   = errors.New("sequentialReaderAt advanced past the requested offset")
	errBadOffset   = errors.New("bad offset value")
	errUnknownSize = errors.New("unknown content size")
)

func (r *fakeReaderAt) SizeAt(off, size int64) (n int64, err error) {
	if size <= 0 {
		return 0, errUnknownSize
	}
	if off < r.pos {
		return 0, errBacktrack
	}
	// extra bytes to accomodate, including 1 extra to detect EOF
	extra := int(off + size + 1 - r.pos - int64(len(r.buf)))
	if extra <= 0 {
		return size, nil
	}
	newBuf := make([]byte, len(r.buf), len(r.buf)+extra)
	copy(newBuf, r.buf)
	r.buf = newBuf
	n1, err1 := io.ReadFull(r.r, r.buf[len(r.buf):cap(r.buf)])
	r.buf = r.buf[:len(r.buf)+n1]
	if n1 == extra { // we read all we asked for
		return r.pos + int64(len(r.buf)) - off - 1, nil // don't report the extra byte in the length
	}
	if err1 == io.ErrUnexpectedEOF {
		err1 = io.EOF
	}
	if n := r.pos + int64(len(r.buf)) - off; n > 0 {
		return n, err1
	}
	return 0, err1
}

func (r *fakeReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n1, err1 := r.SizeAt(off, int64(len(p)))
	if n1 > 0 {
		copy(p[0:n1], r.buf[off-r.pos:])
	}
	return int(n1), err1
}

func (r *fakeReaderAt) ReleaseAt(off int64) error {
	if off < r.pos {
		return errBacktrack
	}
	if off < r.pos+int64(len(r.buf)) { // offset falls within the buffer
		r.buf = r.buf[off-r.pos:]
	} else {
		r.buf = nil
	}
	r.pos = off
	return nil
}
