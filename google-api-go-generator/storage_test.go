package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"
	storage "google.golang.org/api/storage/v1"
)

type myHandler struct {
	location string
	r        *http.Request
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r = r
	if h.location != "" {
		w.Header().Set("Location", h.location)
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

	f := bytes.NewBufferString("fake media data")
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
}

func TestResumableMedia(t *testing.T) {
	handler := &myHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	handler.location = server.URL
	client := &http.Client{}
	s, err := storage.New(client)
	if err != nil {
		t.Fatalf("unable to create service: %v", err)
	}
	s.BasePath = server.URL

	const data = "fake resumable media data"
	f := strings.NewReader(data)
	o := &storage.Object{
		Bucket:          "mybucket",
		Name:            "filename",
		ContentType:     "plain/text",
		ContentEncoding: "utf-8",
		ContentLanguage: "en",
	}
	_, err = s.Objects.Insert("mybucket", o).Name("filename").ResumableMedia(context.Background(), f, int64(len(data)), "text/plain").Do()
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
	if k := "Content-Type"; len(g.Header[k]) != 0 {
		t.Errorf("header %q = %#v; want nil", k, g.Header[k])
	}
	if w, k := "gzip", "Accept-Encoding"; len(g.Header[k]) != 1 || g.Header[k][0] != w {
		t.Errorf("header %q = %#v; want %q", k, g.Header[k], w)
	}
	if w := int64(0); g.ContentLength != w {
		t.Errorf("ContentLength = %v; want %v", g.ContentLength, w)
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
}
