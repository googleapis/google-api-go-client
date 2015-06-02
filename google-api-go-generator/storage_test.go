package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"

	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	storage "google.golang.org/api/storage/v1"
)

type briefString string

func (s briefString) String() string {
	if len(s) < 20 {
		return string(s)
	}
	return fmt.Sprintf("%v...%d chars...%v", s[0:8], len(s)-16, s[len(s)-8:])
}

type myResp struct {
	statusCode int
	header     http.Header
	body       []byte
}

type myHandler struct {
	resps  []myResp
	r      *http.Request
	reqs   []*http.Request
	body   []byte
	bodies [][]byte
	err    error
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r = r
	h.reqs = append(h.reqs, r)
	h.body, h.err = ioutil.ReadAll(r.Body)
	h.bodies = append(h.bodies, h.body)
	if len(h.resps) > 0 {
		resp := h.resps[0]
		h.resps = h.resps[1:]
		for k, values := range resp.header {
			for _, value := range values {
				w.Header().Add(k, value)
			}
		}
		if resp.statusCode > 0 {
			w.WriteHeader(resp.statusCode)
		}
		if resp.body != nil {
			io.Copy(w, bytes.NewReader(resp.body))
		}
		return
	}
	fmt.Fprintf(w, "{}")
}

func TestMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL

	const body = "fake media data"
	f := strings.NewReader(body)
	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Media(f).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w := "POST"; g.Method != w {
		t.Errorf("Method = %q; want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("Proto = %q; want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("ProtoMajor = %v; want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("ProtoMinor = %v; want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w, k := "multipart/related; boundary=", "Content-Type"; len(g.Header[k]) != 1 || !strings.HasPrefix(g.Header[k][0], w) {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w := int64(-1); g.ContentLength != w {
		t.Errorf("ContentLength = %v; want %v", g.ContentLength, w)
	}
	if w := "chunked"; len(g.TransferEncoding) != 1 || g.TransferEncoding[0] != w {
		t.Errorf("TransferEncoding = %#v; want %q", g.TransferEncoding, w)
	}
	if w := strings.TrimPrefix(s.BasePath, "http://"); g.Host != w {
		t.Errorf("Host = %q; want %q", g.Host, w)
	}
	if g.Form != nil {
		t.Errorf("Form = %#v; want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("PostForm = %#v; want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("MultipartForm = %#v; want nil", g.MultipartForm)
	}
	if w := s.BasePath + "/b/mybucket/o?alt=json&uploadType=multipart"; g.RequestURI != w {
		t.Errorf("RequestURI = %q; want %q", g.RequestURI, w)
	}
	if w := "\r\n\r\n" + body + "\r\n"; !strings.Contains(string(handler.body), w) {
		t.Errorf("Body = %q, want substring %q", handler.body, w)
	}
	if handler.err != nil {
		t.Errorf("handler err = %v, want nil", handler.err)
	}
}

// verifyResumableMediaRequest verifies resumable upload request headers and body
func verifyResumableMediaRequest(t *testing.T, g *http.Request, body []byte, contentType string, contentRange string, contentLength int64, bodyWanted string) {
	if w := "POST"; g.Method != w {
		t.Errorf("Method = %q; want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("Proto = %q; want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("ProtoMajor = %v; want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("ProtoMinor = %v; want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w, k := contentType, "Content-Type"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %v", k, g.Header[k], w)
	}
	if w, k := contentRange, "Content-Range"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %v", k, g.Header[k], w)
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w := contentLength; g.ContentLength != w {
		t.Errorf("ContentLength = %v; want %v", g.ContentLength, w)
	}
	if s, w := string(body), bodyWanted; s != w {
		t.Errorf("body = %q; want %q", briefString(s), briefString(w))
	}
	if len(g.TransferEncoding) != 0 {
		t.Errorf("TransferEncoding = %#v; want nil", g.TransferEncoding)
	}
	if g.Form != nil {
		t.Errorf("Form = %#v; want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("PostForm = %#v; want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("MultipartForm = %#v; want nil", g.MultipartForm)
	}
}

// testResumableMediaHelper creates a httptest.Server and runs resumable upload flow against it.
func testResumableMediaHelper(t *testing.T, received int, rdr io.Reader, opt ...googleapi.UploadOption) *myHandler {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	handler.resps = []myResp{
		// response to session initiation request: empty 200 OK with Location header
		myResp{header: http.Header{"Location": []string{server.URL}}},
		// response to first chunk: empty 500 to trigger status request
		myResp{statusCode: 500},
		// response to status request: 308 Resume Incomplete with Range header
		myResp{statusCode: 308, header: http.Header{"Range": []string{fmt.Sprintf("bytes=0-%v", received-1)}}},
		// response to the resumed content: 200 OK, meaning upload complete
		myResp{body: []byte("{}")},
	}

	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL

	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}

	// This should trigger the HTTP request/response sequence above to play out. The final
	// request to the server should be the resumed upload of the partial content (bytes
	// 6-end).
	_, err = s.Objects.Insert("mybucket", o).Name("filename").ResumableMedia(context.Background(), rdr, opt...).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	if handler.err != nil {
		t.Errorf("handler err = %v, want nil", handler.err)
	}
	if l, w := len(handler.reqs), 4; l != w {
		t.Fatalf("Seen %v requests; want %v", l, w)
	}
	return handler
}

// sizedReader wraps strings.Reader with an additional Size() method.
type sizedReader struct {
	strings.Reader
}

func (sr *sizedReader) Size() int64 { return int64(sr.Len()) }

func TestResumableMedia(t *testing.T) {
	const received = 6 // pretend server received this many bytes before server error
	data := strings.Repeat("fake resumable media data", 2000)
	dataLen := int64(len(data))

	// Temporarily shorten the default chunk size so we can observe chunking.
	// We've chosen new default chunkSize (dataLen-3) to be shorter than the full content
	// to be sent (dataLen), but longer than the rest of the content to be sent after upload
	// resumption (dataLen-6).
	// That way the upload should complete in the same number of requests regardless of
	// whether chunking is taking place, so we can use the same sequence of fake server
	// responses in all cases. We observe chunking indirectly, by observing content
	// size and content range header.
	chunkSize, origChunkSize := dataLen-3, googleapi.ChunkSize
	googleapi.ChunkSize = chunkSize
	defer func() { googleapi.ChunkSize = origChunkSize }()

	// Note that strings.Reader supports io.ReaderAt interface, but content will be
	// buffered since content size is not known. We expect chunking to take place.
	t.Log("Resumable upload, buffered chunker")
	handler := testResumableMediaHelper(t, received, strings.NewReader(data))
	// Verify first chunk with data payload
	verifyResumableMediaRequest(t, handler.reqs[1], handler.bodies[1],
		"text/plain; charset=utf-8",
		fmt.Sprintf("bytes %v-%v/%v", 0, chunkSize-1, "*"),
		chunkSize,
		data[:chunkSize])
	// Verify second chunk with data payload (one after upload resumption)
	verifyResumableMediaRequest(t, handler.r, handler.body,
		"text/plain; charset=utf-8",
		fmt.Sprintf("bytes %v-%v/%v", received, dataLen-1, dataLen),
		dataLen-received,
		data[received:])

	// Here we explicitly specify content size, allowing io.ReaderAt usage. No chunking expected.
	t.Log("Resumable upload, io.ReaderAt, no chunking")
	sr := &sizedReader{*strings.NewReader(data)}
	handler = testResumableMediaHelper(t, received, sr, googleapi.SetMediaType("text/html; charset=utf-8"))
	verifyResumableMediaRequest(t, handler.reqs[1], handler.bodies[1],
		"text/html; charset=utf-8",
		fmt.Sprintf("bytes %v-%v/%v", 0, dataLen-1, dataLen),
		dataLen,
		data)
	verifyResumableMediaRequest(t, handler.r, handler.body,
		"text/html; charset=utf-8",
		fmt.Sprintf("bytes %v-%v/%v", received, dataLen-1, dataLen),
		dataLen-received,
		data[received:])
}

func TestNoMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL

	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w := "POST"; g.Method != w {
		t.Errorf("Method = %q; want %q", g.Method, w)
	}
	if w := "HTTP/1.1"; g.Proto != w {
		t.Errorf("Proto = %q; want %q", g.Proto, w)
	}
	if w := 1; g.ProtoMajor != w {
		t.Errorf("ProtoMajor = %v; want %v", g.ProtoMajor, w)
	}
	if w := 1; g.ProtoMinor != w {
		t.Errorf("ProtoMinor = %v; want %v", g.ProtoMinor, w)
	}
	if w, k := "google-api-go-client/0.5", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w, k := "application/json", "Content-Type"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w := int64(116); g.ContentLength != w {
		t.Errorf("ContentLength = %v; want %v", g.ContentLength, w)
	}
	if len(g.TransferEncoding) != 0 {
		t.Errorf("TransferEncoding = %#v; want []string{}", g.TransferEncoding)
	}
	if w := strings.TrimPrefix(s.BasePath, "http://"); g.Host != w {
		t.Errorf("Host = %q; want %q", g.Host, w)
	}
	if g.Form != nil {
		t.Errorf("Form = %#v; want nil", g.Form)
	}
	if g.PostForm != nil {
		t.Errorf("PostForm = %#v; want nil", g.PostForm)
	}
	if g.MultipartForm != nil {
		t.Errorf("MultipartForm = %#v; want nil", g.MultipartForm)
	}
	if w := s.BasePath + "/b/mybucket/o?alt=json"; g.RequestURI != w {
		t.Errorf("RequestURI = %q; want %q", g.RequestURI, w)
	}
	if w := `{"bucket":"mybucket","contentEncoding":"utf-8","contentLanguage":"en","contentType":"plain/text","name":"filename"}` + "\n"; string(handler.body) != w {
		t.Errorf("Body = %q, want %q", handler.body, w)
	}
	if handler.err != nil {
		t.Errorf("handler err = %v, want nil", handler.err)
	}
}

func TestMediaErrHandling(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL

	const body = "fake media data"
	f := strings.NewReader(body)
	// The combination of TimeoutReader and OneByteReader causes the first byte to
	// be successfully delivered, but then a timeout error is reported.  This
	// allows us to test the goroutine within the getMediaType function.
	r := iotest.TimeoutReader(iotest.OneByteReader(f))
	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Media(r).Do()
	if err == nil || !strings.Contains(err.Error(), "timeout") {
		t.Errorf("expected timeout error, got %v", err)
	}
	if handler.err != nil {
		t.Errorf("handler err = %v, want nil", handler.err)
	}
}

func TestUserAgent(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL
	s.UserAgent = "myagent/1.0"

	f := strings.NewReader("fake media data")
	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Media(f).Do()
	if err != nil {
		t.Fatalf("unable to insert object: %v", err)
	}
	g := handler.r
	if w, k := "google-api-go-client/0.5 myagent/1.0", "User-Agent"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
}
