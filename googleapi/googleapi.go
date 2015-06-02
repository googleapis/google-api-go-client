// Copyright 2011 Google Inc. All rights reserved.
// Copyright (C) 2015 Motorola Mobility LLC All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package googleapi contains the common code shared by all Google API
// libraries.
package googleapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/api/googleapi/internal/uritemplates"
)

// ContentTyper is an interface for Readers which know (or would like
// to override) their Content-Type. If a media body doesn't implement
// ContentTyper, the type is sniffed from the content using
// http.DetectContentType.
type ContentTyper interface {
	ContentType() string
}

// Reader is the type accepted by ResumableMedia uploads.
// It can be any io.Reader, or as a special case for extra configuration
// it may be a *UploadParameters.
type Reader interface {
	io.Reader
}
type UploadParameters struct {
	// Reader is the source to read from.
	// It is the only required field.
	io.Reader
	// Size optionally lists the Reader size in bytes.
	// It may be set to 0 if unknown.
	Size int64
	// ChunkSize optionally specifies the size of upload chunk.
	// If zero, chunk size will be auto-selected, or chunking may be skipped altogether
	// if not required (when content size is known up-front and reader supports
	// io.ReaderAt interface). ChunkSize that is not a multiple of 256KiB (256<<10 bytes) will be
	// automatically upgraded to the next smallest multiple of 256KiB, in accordance with
	// the Google Cloud Storage specifications.
	// Note the memory impact of setting chunk size: one full chunk will be buffered in memory.
	ChunkSize int64
	// MediaType optionally specifies content MIME type, such as "image/png".
	// If empty, it is auto-detected using http.DetectContentType.
	MediaType string
}

const (
	Version = "0.5"

	// statusResumeIncomplete is the code returned by the Google uploader when the transfer is not yet complete.
	statusResumeIncomplete = 308

	// UserAgent is the header string used to identify this package.
	UserAgent = "google-api-go-client/" + Version

	// uploadPause determines the delay between failed upload attempts
	uploadPause = 1 * time.Second
)

// Error contains an error response from the server.
type Error struct {
	// Code is the HTTP response status code and will always be populated.
	Code int `json:"code"`
	// Message is the server response message and is only populated when
	// explicitly referenced by the JSON server response.
	Message string `json:"message"`
	// Body is the raw response returned by the server.
	// It is often but not always JSON, depending on how the request fails.
	Body string

	Errors []ErrorItem
}

// ErrorItem is a detailed error code & message from the Google API frontend.
type ErrorItem struct {
	// Reason is the typed error code. For example: "some_example".
	Reason string `json:"reason"`
	// Message is the human-readable description of the error.
	Message string `json:"message"`
}

func (e *Error) Error() string {
	if len(e.Errors) == 0 && e.Message == "" {
		return fmt.Sprintf("googleapi: got HTTP response code %d with body: %v", e.Code, e.Body)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "googleapi: Error %d: ", e.Code)
	if e.Message != "" {
		fmt.Fprintf(&buf, "%s", e.Message)
	}
	if len(e.Errors) == 0 {
		return strings.TrimSpace(buf.String())
	}
	if len(e.Errors) == 1 && e.Errors[0].Message == e.Message {
		fmt.Fprintf(&buf, ", %s", e.Errors[0].Reason)
		return buf.String()
	}
	fmt.Fprintln(&buf, "\nMore details:")
	for _, v := range e.Errors {
		fmt.Fprintf(&buf, "Reason: %s, Message: %s\n", v.Reason, v.Message)
	}
	return buf.String()
}

type errorReply struct {
	Error *Error `json:"error"`
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err == nil {
		jerr := new(errorReply)
		err = json.Unmarshal(slurp, jerr)
		if err == nil && jerr.Error != nil {
			if jerr.Error.Code == 0 {
				jerr.Error.Code = res.StatusCode
			}
			jerr.Error.Body = string(slurp)
			return jerr.Error
		}
	}
	return &Error{
		Code: res.StatusCode,
		Body: string(slurp),
	}
}

type MarshalStyle bool

var WithDataWrapper = MarshalStyle(true)
var WithoutDataWrapper = MarshalStyle(false)

func (wrap MarshalStyle) JSONReader(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if wrap {
		buf.Write([]byte(`{"data": `))
	}
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	if wrap {
		buf.Write([]byte(`}`))
	}
	return buf, nil
}

func getMediaType(media io.Reader) (io.Reader, string) {
	if typer, ok := media.(ContentTyper); ok {
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
	if typer, ok := media.(ContentTyper); ok {
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

type Lengther interface {
	Len() int
}

// endingWithErrorReader from r until it returns an error.  If the
// final error from r is io.EOF and e is non-nil, e is used instead.
type endingWithErrorReader struct {
	r io.Reader
	e error
}

func (er endingWithErrorReader) Read(p []byte) (n int, err error) {
	n, err = er.r.Read(p)
	if err == io.EOF && er.e != nil {
		err = er.e
	}
	return
}

func typeHeader(contentType string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", contentType)
	return h
}

// countingWriter counts the number of bytes it receives to write, but
// discards them.
type countingWriter struct {
	n *int64
}

func (w countingWriter) Write(p []byte) (int, error) {
	*w.n += int64(len(p))
	return len(p), nil
}

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
// The return value is the content-length of the entire multpart body.
func ConditionallyIncludeMedia(media io.Reader, bodyp *io.Reader, ctypep *string) (cancel func(), ok bool) {
	if media == nil {
		return
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
	cancel = func() { pw.CloseWithError(errAborted) }
	return cancel, true
}

var errAborted = errors.New("googleapi: upload aborted")

// ProgressUpdater is a function that is called upon every progress update of a resumable upload.
// This is the only part of a resumable upload (from googleapi) that is usable by the developer.
// The remaining usable pieces of resumable uploads is exposed in each auto-generated API.
type ProgressUpdater func(current, total int64)

// ResumableUpload is used by the generated APIs to provide resumable uploads.
// It is not used by developers directly.
type ResumableUpload struct {
	Client *http.Client
	// URI is the resumable resource destination provided by the server after specifying "&uploadType=resumable".
	URI       string
	UserAgent string // User-Agent for header of the request
	// Media is the object being uploaded.
	Media io.ReaderAt
	// MediaType defines the media type, e.g. "image/jpeg".
	MediaType string
	// ContentLength is the full size of the object being uploaded.
	ContentLength int64
	// ChunkSize is the chunk size to be using during upload. Zero or negative value means use default.
	ChunkSize int64

	mu       sync.Mutex // guards progress
	progress int64      // number of bytes uploaded so far
	started  bool       // whether the upload has been started

	// Callback is an optional function that will be called upon every progress update.
	Callback ProgressUpdater
}

var (
	// rangeRE matches the transfer status response from the server. $1 is the last byte index uploaded.
	rangeRE = regexp.MustCompile(`^bytes=0\-(\d+)$`)
	// ChunkSize is the size of the chunks created during a resumable upload and should be a multiple of
	// 256KB, per Google Cloud Storage requirements. While Google Cloud Storage does not specify maximum,
	// note that classic app engine urlfetch limits outgoing HTTP request size to 10MB max.
	ChunkSize int64 = 8 << 20
)

// Configure interrogates the supplied io.Reader for additional capabilities, and initializes
// the following fields in the ResumableUpload struct:
//  - Media is set to an appropriate io.ReaderAt, which is possibly a *fakeReaderAt.
//  - MediaType is set to user-provided value, or guessed using http.DetectContentType.
//  - ContentLength is set to user-provided value, or possibly guessed, or left as 0
//    indicating that content length is unknown.
//  - ChunkSize is set to user-provided value, or to package default value,
//    or left as 0 indicating that chunking should be skipped.
//
// The method expects all the above fields to be at their zero values prior to the method
// invocation. No other fields will be touched by this method. The method must be invoked
// prior to starting the resumable upload.
//
// The additional io.Reader capabilities that this method looks for are:
//  - whether reader supports io.ReaderAt interface
//  - whether reader supports os.File.Stat() method
//  - whether reader is also an UploadParameters struct value or struct pointer
//  - whether Reader embedded in UploadParameters supports io.ReaderAt and os.File.Stat()
//
// The ultimate goal of this configuration is to ensure that resumable upload is done
// in the most efficient manner whenever possible, while falling back to the less efficient
// method when required (with buffering content and chunking). For example, the most
// common case of user passing *os.File as io.Reader is handled efficiently - buffering
// and chunking are avoided.
// This method is called from, and private to the auto-generated API code, and should
// not be invoked directly by the user.
func (rx *ResumableUpload) Configure(r io.Reader) {
	type fileLike interface { // relevant subset of os.File methods
		io.ReaderAt
		Stat() (os.FileInfo, error)
	}
	if file, ok := r.(fileLike); ok {
		if fileinfo, err := file.Stat(); err == nil {
			rx.ContentLength = fileinfo.Size()
			rx.Media = file
			rx.MediaType = DetectMediaType(file)
			return
		}
	}
	// Check for UploadParameters. Allow either struct or *struct since both
	// forms satisfy io.Reader and either could be used by the user.
	params, ok := r.(*UploadParameters)
	if !ok {
		if params1, ok := r.(UploadParameters); ok {
			params = &params1
		}
	}
	if params != nil {
		rx.ContentLength = params.Size
		rx.MediaType = params.MediaType
		rx.ChunkSize = calcChunkSize(params.ChunkSize)
		if rx.Media, ok = params.Reader.(io.ReaderAt); ok {
			if params.Size <= 0 {
				// try to guess size from File
				rx.Media = nil
				if file, ok := params.Reader.(fileLike); ok {
					if fileinfo, err := file.Stat(); err == nil {
						rx.ContentLength = fileinfo.Size()
						rx.Media = file
					}
				}
			}
		}
	}
	if rx.Media == nil {
		rx.Media = &fakeReaderAt{r: r}
		// always force chunking when using fakeReaderAt
		if rx.ChunkSize <= 0 {
			rx.ChunkSize = ChunkSize // use package default
		}
	}
	if rx.MediaType == "" {
		rx.MediaType = DetectMediaType(rx.Media)
	}
}

// Progress returns the number of bytes uploaded at this point.
func (rx *ResumableUpload) Progress() int64 {
	rx.mu.Lock()
	defer rx.mu.Unlock()
	return rx.progress
}

func (rx *ResumableUpload) transferStatus() (int64, *http.Response, error) {
	req, _ := http.NewRequest("POST", rx.URI, nil)
	req.ContentLength = 0
	req.Header.Set("User-Agent", rx.UserAgent)
	req.Header.Set("Content-Range", "bytes */*")
	res, err := rx.Client.Do(req)
	if err != nil || res.StatusCode != statusResumeIncomplete {
		return 0, res, err
	}
	var start int64
	if m := rangeRE.FindStringSubmatch(res.Header.Get("Range")); len(m) == 2 {
		start, err = strconv.ParseInt(m[1], 10, 64)
		if err != nil {
			return 0, nil, fmt.Errorf("unable to parse range size %v", m[1])
		}
		start += 1 // Start at the next byte
	}
	return start, res, nil
}

type chunk struct {
	body io.Reader
	size int64
	err  error
}

func (rx *ResumableUpload) transferChunks(ctx context.Context) (*http.Response, error) {
	var start int64
	var err error
	res := &http.Response{}
	if rx.started {
		start, res, err = rx.transferStatus()
		if err != nil || res.StatusCode != statusResumeIncomplete {
			return res, err
		}
	}
	rx.started = true

	for {
		select { // Check for cancellation
		case <-ctx.Done():
			res.StatusCode = http.StatusRequestTimeout
			return res, ctx.Err()
		default:
		}

		var reqSize int64
		if fr, ok := rx.Media.(*fakeReaderAt); ok && rx.ContentLength <= 0 {
			chunkSize_ := rx.ChunkSize
			if chunkSize_ <= 0 {
				chunkSize_ = ChunkSize // use package default
			}
			size, err := fr.SizeAt(start, int(chunkSize_))
			reqSize = int64(size)
			if err != nil && err != io.EOF {
				return res, err
			}
			if err == io.EOF { // reached end of stream, this is our final chunk
				rx.ContentLength = start + int64(reqSize) // finally learned content length
			}
		} else { // we have real ReaderAt with known content size
			reqSize = rx.ContentLength - start
			// use chunking only if explicitly requested
			if rx.ChunkSize > 0 && reqSize > rx.ChunkSize {
				reqSize = rx.ChunkSize
			}
		}
		r := io.NewSectionReader(rx.Media, start, reqSize)
		req, _ := http.NewRequest("POST", rx.URI, r)
		req.ContentLength = reqSize
		if rx.ContentLength <= 0 {
			req.Header.Set("Content-Range", fmt.Sprintf("bytes %v-%v/*", start, start+reqSize-1))
		} else {
			req.Header.Set("Content-Range", fmt.Sprintf("bytes %v-%v/%v", start, start+reqSize-1, rx.ContentLength))
		}
		req.Header.Set("Content-Type", rx.MediaType)
		req.Header.Set("User-Agent", rx.UserAgent)
		res, err = rx.Client.Do(req)
		start += reqSize
		if err == nil && (res.StatusCode == statusResumeIncomplete || res.StatusCode == http.StatusOK) {
			rx.mu.Lock()
			rx.progress = start // keep track of number of bytes sent so far
			rx.mu.Unlock()
			if rx.Callback != nil {
				rx.Callback(start, rx.ContentLength)
			}
		}
		if err != nil || res.StatusCode != statusResumeIncomplete {
			break
		}
	}
	return res, err
}

// calcChunkSize returns smallest multiple of 256K (GCS requirement) greater or equal
// to chunkSize.  However, 0 or negative values are left unmodified.
func calcChunkSize(size int64) int64 {
	const (
		c     = int64(256 << 10)
		cmask = c - 1
	)
	if size <= 0 || size&cmask == 0 {
		return size
	}
	return ((size >> 18) + 1) << 18
}

var sleep = time.Sleep // override in unit tests

// Upload starts the process of a resumable upload with a cancellable context.
// It retries indefinitely (with a pause of uploadPause between attempts) until cancelled.
// It is called from the auto-generated API code and is not visible to the user.
// rx is private to the auto-generated API code.
func (rx *ResumableUpload) Upload(ctx context.Context) (*http.Response, error) {
	var res *http.Response
	var err error
	for {
		res, err = rx.transferChunks(ctx)
		if err != nil || res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
			return res, err
		}
		select { // Check for cancellation
		case <-ctx.Done():
			res.StatusCode = http.StatusRequestTimeout
			return res, ctx.Err()
		default:
		}
		sleep(uploadPause)
	}
	return res, err
}

func ResolveRelative(basestr, relstr string) string {
	u, _ := url.Parse(basestr)
	rel, _ := url.Parse(relstr)
	u = u.ResolveReference(rel)
	us := u.String()
	us = strings.Replace(us, "%7B", "{", -1)
	us = strings.Replace(us, "%7D", "}", -1)
	return us
}

// has4860Fix is whether this Go environment contains the fix for
// http://golang.org/issue/4860
var has4860Fix bool

// init initializes has4860Fix by checking the behavior of the net/http package.
func init() {
	r := http.Request{
		URL: &url.URL{
			Scheme: "http",
			Opaque: "//opaque",
		},
	}
	b := &bytes.Buffer{}
	r.Write(b)
	has4860Fix = bytes.HasPrefix(b.Bytes(), []byte("GET http"))
}

// SetOpaque sets u.Opaque from u.Path such that HTTP requests to it
// don't alter any hex-escaped characters in u.Path.
func SetOpaque(u *url.URL) {
	u.Opaque = "//" + u.Host + u.Path
	if !has4860Fix {
		u.Opaque = u.Scheme + ":" + u.Opaque
	}
}

// Expand subsitutes any {encoded} strings in the URL passed in using
// the map supplied.
//
// This calls SetOpaque to avoid encoding of the parameters in the URL path.
func Expand(u *url.URL, expansions map[string]string) {
	expanded, err := uritemplates.Expand(u.Path, expansions)
	if err == nil {
		u.Path = expanded
		SetOpaque(u)
	}
}

// CloseBody is used to close res.Body.
// Prior to calling Close, it also tries to Read a small amount to see an EOF.
// Not seeing an EOF can prevent HTTP Transports from reusing connections.
func CloseBody(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}
	// Justification for 3 byte reads: two for up to "\r\n" after
	// a JSON/XML document, and then 1 to see EOF if we haven't yet.
	// TODO(bradfitz): detect Go 1.3+ and skip these reads.
	// See https://codereview.appspot.com/58240043
	// and https://codereview.appspot.com/49570044
	buf := make([]byte, 1)
	for i := 0; i < 3; i++ {
		_, err := res.Body.Read(buf)
		if err != nil {
			break
		}
	}
	res.Body.Close()

}

// VariantType returns the type name of the given variant.
// If the map doesn't contain the named key or the value is not a []interface{}, "" is returned.
// This is used to support "variant" APIs that can return one of a number of different types.
func VariantType(t map[string]interface{}) string {
	s, _ := t["type"].(string)
	return s
}

// ConvertVariant uses the JSON encoder/decoder to fill in the struct 'dst' with the fields found in variant 'v'.
// This is used to support "variant" APIs that can return one of a number of different types.
// It reports whether the conversion was successful.
func ConvertVariant(v map[string]interface{}, dst interface{}) bool {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return false
	}
	return json.Unmarshal(buf.Bytes(), dst) == nil
}

// A Field names a field to be retrieved with a partial response.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
//
// Partial responses can dramatically reduce the amount of data that must be sent to your application.
// In order to request partial responses, you can specify the full list of fields
// that your application needs by adding the Fields option to your request.
//
// Field strings use camelCase with leading lower-case characters to identify fields within the response.
//
// For example, if your response has a "NextPageToken" and a slice of "Items" with "Id" fields,
// you could request just those fields like this:
//
//     svc.Events.List().Fields("nextPageToken", "items/id").Do()
//
// or if you were also interested in each Item's "Updated" field, you can combine them like this:
//
//     svc.Events.List().Fields("nextPageToken", "items(id,updated)").Do()
//
// More information about field formatting can be found here:
// https://developers.google.com/+/api/#fields-syntax
//
// Another way to find field names is through the Google API explorer:
// https://developers.google.com/apis-explorer/#p/
type Field string

// CombineFields combines fields into a single string.
func CombineFields(s []Field) string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = string(v)
	}
	return strings.Join(r, ",")
}
