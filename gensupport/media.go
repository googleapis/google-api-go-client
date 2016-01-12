// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"google.golang.org/api/googleapi"
)

const sniffBuffSize = 512

func NewContentSniffer(r io.Reader) *ContentSniffer {
	return &ContentSniffer{r: r}
}

// ContentSniffer wraps a Reader, and reports the content type determined by sniffing up to 512 bytes from the Reader.
type ContentSniffer struct {
	r     io.Reader
	start []byte // buffer for the sniffed bytes.
	err   error  // set to any error encountered while reading bytes to be sniffed.

	ctype   string // set on first sniff.
	sniffed bool   // set to true on first sniff.
}

func (cs *ContentSniffer) Read(p []byte) (n int, err error) {
	// Ensure that the content type is sniffed before any data is consumed from Reader.
	_, _ = cs.ContentType()

	if len(cs.start) > 0 {
		n := copy(p, cs.start)
		cs.start = cs.start[n:]
		return n, nil
	}

	// We may have read some bytes into start while sniffing, even if the read ended in an error.
	// We should first return those bytes, then the error.
	if cs.err != nil {
		return 0, cs.err
	}

	// Now we have handled all bytes that were buffered while sniffing.  Now just delegate to the underlying reader.
	return cs.r.Read(p)
}

// ContentType returns the sniffed content type, and whether the content type was succesfully sniffed.
func (cs *ContentSniffer) ContentType() (string, bool) {
	if cs.sniffed {
		return cs.ctype, cs.ctype != ""
	}
	cs.sniffed = true
	// If ReadAll hits EOF, it returns err==nil.
	cs.start, cs.err = ioutil.ReadAll(io.LimitReader(cs.r, sniffBuffSize))

	// Don't try to detect the content type based on possibly incomplete data.
	if cs.err != nil {
		return "", false
	}

	cs.ctype = http.DetectContentType(cs.start)
	return cs.ctype, true
}

// DetectContentType determines the content type of the supplied reader.
// If the content type is already known, it can be specified via ctype.
// Otherwise, the content of media will be sniffed to determine the content type.
// If media implements googleapi.ContentTyper (deprecated), this will be used
// instead of sniffing the content.
// After calling DetectContentType the caller must not perform further reads on
// media, but rather read from the Reader that is returned.
func DetectContentType(media io.Reader, ctype string) (r io.Reader, ctype string) {
	// Note: callers could avoid calling DetectContentType if ctype != "",
	// but doing the check inside this function reduces the amount of
	// generated code.
	if ctype != "" {
		return media, ctype
	}

	// For backwards compatability, allow clients to set content
	// type by providing a ContentTyper for media.
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return media, typer.ContentType()
	}

	sniffer := gensupport.NewContentSniffer(r)
	if ctype, ok := sniffer.ContentType(); ok {
		return sniffer, ctype
	}
	// If content type could not be sniffed, reads from sniffer will eventually fail with an error.
	return sniffer, ""
}

// IncludeMedia combines an existing HTTP body with media content to create a multipart/related HTTP body.
//
// bodyp is an in/out parameter.  It should initially point to the
// reader of the application/json (or whatever) payload to send in the
// API request.  It's updated to point to the multipart body reader.
//
// ctypep is an in/out parameter.  It should initially point to the
// content type of the bodyp, usually "application/json".  It's updated
// to the "multipart/related" content type, with random boundary.
//
// The return value is a function that can be used to close the bodyp Reader with an error.
func IncludeMedia(media io.Reader, mediaType string, bodyp *io.Reader, ctypep *string) func() {
	body, bodyType := *bodyp, *ctypep

	pr, pw := io.Pipe()
	mpw := multipart.NewWriter(pw)
	*bodyp = pr
	*ctypep = "multipart/related; boundary=" + mpw.Boundary()
	go func() {
		w, err := mpw.CreatePart(typeHeader(bodyType))
		if err != nil {
			mpw.Close()
			pw.CloseWithError(fmt.Errorf("googleapi: body CreatePart failed: %v", err))
			return
		}
		_, err = io.Copy(w, body)
		if err != nil {
			mpw.Close()
			pw.CloseWithError(fmt.Errorf("googleapi: body Copy failed: %v", err))
			return
		}

		w, err = mpw.CreatePart(typeHeader(mediaType))
		if err != nil {
			mpw.Close()
			pw.CloseWithError(fmt.Errorf("googleapi: media CreatePart failed: %v", err))
			return
		}
		_, err = io.Copy(w, media)
		if err != nil {
			mpw.Close()
			pw.CloseWithError(fmt.Errorf("googleapi: media Copy failed: %v", err))
			return
		}
		mpw.Close()
		pw.Close()
	}()
	return func() { pw.CloseWithError(errAborted) }
}

var errAborted = errors.New("googleapi: upload aborted")

func getMediaType(media io.Reader) (io.Reader, string) {
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return media, typer.ContentType()
	}

	sniffer := NewContentSniffer(media)
	typ, ok := sniffer.ContentType()
	if !ok {
		// TODO(mcgreevy): Remove this default.  It maintains the semantics of the existing code,
		// but should not be relied on.
		typ = "application/octet-stream"
	}
	return sniffer, typ
}

// DetectMediaType detects and returns the content type of the provided media.
// If the type can not be determined, "application/octet-stream" is returned.
func DetectMediaType(media io.ReaderAt) string {
	if typer, ok := media.(googleapi.ContentTyper); ok {
		return typer.ContentType()
	}

	typ := "application/octet-stream"
	buf := make([]byte, 1024)
	n, err := media.ReadAt(buf, 0)
	buf = buf[:n]
	if err == nil || err == io.EOF {
		typ = http.DetectContentType(buf)
	}
	return typ
}

func typeHeader(contentType string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", contentType)
	return h
}
