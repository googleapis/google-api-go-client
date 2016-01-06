// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"google.golang.org/api/googleapi"
)

// ConditionallyIncludeMedia does nothing if media is nil.
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
func ConditionallyIncludeMedia(media io.Reader, bodyp *io.Reader, ctypep *string) func() {
	if media == nil {
		return func() {}
	}
	// Get the media type, which might return a different reader instance.
	var mediaType string
	media, mediaType = getMediaType(media)

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

	pr, pw := io.Pipe()
	typ := "application/octet-stream"
	buf, err := ioutil.ReadAll(io.LimitReader(media, 512))
	if err != nil {
		pw.CloseWithError(fmt.Errorf("error reading media: %v", err))
		return pr, typ
	}
	typ = http.DetectContentType(buf)
	mr := io.MultiReader(bytes.NewReader(buf), media)
	go func() {
		_, err = io.Copy(pw, mr)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("error reading media: %v", err))
			return
		}
		pw.Close()
	}()
	return pr, typ
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
