// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

// reqTmpl is used to match the output of writing an HTTP request.
const reqTmpl = "GET %v HTTP/1.1\r\nHost: %v\r\nUser-Agent: %v\r\n\r\n"

type SetOpaqueTest struct {
	in             *url.URL
	wantRequestURI string
}

var setOpaqueTests = []SetOpaqueTest{
	// no path
	{
		&url.URL{
			Scheme: "http",
			Host:   "www.golang.org",
		},
		"http://www.golang.org",
	},
	// path
	{
		&url.URL{
			Scheme: "http",
			Host:   "www.golang.org",
			Path:   "/",
		},
		"http://www.golang.org/",
	},
	// file with hex escaping
	{
		&url.URL{
			Scheme: "https",
			Host:   "www.golang.org",
			Path:   "/file%20one&two",
		},
		"https://www.golang.org/file%20one&two",
	},
	// query
	{
		&url.URL{
			Scheme:   "http",
			Host:     "www.golang.org",
			Path:     "/",
			RawQuery: "q=go+language",
		},
		"http://www.golang.org/?q=go+language",
	},
	// file with hex escaping in path plus query
	{
		&url.URL{
			Scheme:   "https",
			Host:     "www.golang.org",
			Path:     "/file%20one&two",
			RawQuery: "q=go+language",
		},
		"https://www.golang.org/file%20one&two?q=go+language",
	},
	// query with hex escaping
	{
		&url.URL{
			Scheme:   "http",
			Host:     "www.golang.org",
			Path:     "/",
			RawQuery: "q=go%20language",
		},
		"http://www.golang.org/?q=go%20language",
	},
}

func TestSetOpaque(t *testing.T) {
	userAgent := "Go http package"
	if isGo11 {
		userAgent = "Go 1.1 package http"
	}

	for _, test := range setOpaqueTests {
		u := *test.in
		SetOpaque(&u)

		w := &bytes.Buffer{}
		r := &http.Request{URL: &u}
		if err := r.Write(w); err != nil {
			t.Errorf("write request: %v", err)
			continue
		}

		exp := fmt.Sprintf(reqTmpl, test.wantRequestURI, test.in.Host, userAgent)
		if got := string(w.Bytes()); got != exp {
			t.Errorf("got %q expected %q", got, exp)
		}
	}
}
