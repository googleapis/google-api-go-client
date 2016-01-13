// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"bytes"
	"io"
)

// ResumableBuffer buffers data from an io.Reader to support uploading media in retryable chunks.
type ResumableBuffer struct {
	media io.Reader

	chunk []byte // The current chunk which is pending upload.  The capacity is the chunk size.
	err   error  // Any error generated when populating chunk by reading media.

	// The absolute position of chunk in the underlying media.
	off int64
}

func NewResumableBuffer(media io.Reader, chunkSize int) *ResumableBuffer {
	return &ResumableBuffer{media: media, chunk: make([]byte, 0, chunkSize)}
}

// Chunk returns the current chunk which is pending upload, the offset in the underlying media from
// which the chunk is drawn, and the size of the chunk.
// To obtain the next chunk, Next must first be called.
func (rb *ResumableBuffer) Chunk() (chunk io.Reader, off int64, size int, err error) {
	// There may already be data in chunk if Next has not been called since the previous call to Chunk.
	if len(rb.chunk) == 0 {
		rb.err = rb.loadChunk()
	}
	return bytes.NewReader(rb.chunk), rb.off, len(rb.chunk), rb.err
}

// loadChunk will read from media into chunk, up to the capacity of chunk.
func (rb *ResumableBuffer) loadChunk() error {
	bufSize := cap(rb.chunk)
	rb.chunk = rb.chunk[:bufSize]

	read := 0
	var err error
	for err == nil && read < bufSize {
		var n int
		n, err = rb.media.Read(rb.chunk[read:])
		read += n
	}
	rb.chunk = rb.chunk[:read]
	return err
}

// Next confirms that the current chunk has been successfully uploaded.
// The next call to Chunk will return the next chunk of data.
// Calls to Next without a corresponding prior call to Chunk will have no effect.
func (rb *ResumableBuffer) Next() {
	rb.off += int64(len(rb.chunk))
	rb.chunk = rb.chunk[0:0]
}
