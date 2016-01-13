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
	start int64
}

func NewResumableBuffer(media io.Reader, chunkSize int) *ResumableBuffer {
	return &ResumableBuffer{media: media, chunk: make([]byte, 0, chunkSize)}
}

// Chunk returns the current chunk which is pending upload, along with its size.
// To obtain the next chunk, Next must first be called.
func (rb *ResumableBuffer) Chunk() (io.Reader, int, error) {
	// There may already be data in chunk if Next has not been called since the previous call to Chunk.
	if len(rb.chunk) == 0 {
		rb.chunk = rb.chunk[:cap(rb.chunk)]
		var n int
		n, rb.err = io.ReadFull(rb.media, rb.chunk)
		if rb.err == io.ErrUnexpectedEOF {
			// It's only unexpected because there was not enough data to fill the buffer.
			// As far as clients of ResumableBuffer are concerned, it's just a plain EOF.
			rb.err = io.EOF
		}
		rb.chunk = rb.chunk[:n]
	}
	return bytes.NewReader(rb.chunk), len(rb.chunk), rb.err
}

// Next confirms that the current chunk has been successfully uploaded.
// The next call to Chunk will return the next chunk of data.
// Calls to Next without a corresponding prior call to Chunk will have no effect.
func (rb *ResumableBuffer) Next() {
	rb.start += int64(len(rb.chunk))
	rb.chunk = rb.chunk[0:0]
}

// Pos returns the number of bytes which preceed the beginning of the current chunk in the supplied media content.
func (rb *ResumableBuffer) Pos() int64 {
	return rb.start
}
