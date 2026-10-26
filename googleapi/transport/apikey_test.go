// Copyright 2026 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transport

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

// TestAPIKeyRedirect checks where the key goes when an http.Client follows
// a redirect through APIKey.
//
// Both test servers listen on the loopback interface, but server A is
// addressed as "localhost" and server B as "127.0.0.1". net/http decides
// whether a redirect leaves the original domain by comparing hostnames, so
// that hop is cross-domain, whereas two httptest servers addressed by their
// URLs are both 127.0.0.1 and count as the same host.
func TestAPIKeyRedirect(t *testing.T) {
	const key = "secret"

	var mu sync.Mutex
	got := map[string]string{} // final path -> "key" query parameter received there
	record := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		got[r.URL.Path] = r.URL.Query().Get("key")
	}
	received := func(path string) string {
		mu.Lock()
		defer mu.Unlock()
		return got[path]
	}

	var addrA string // server A, addressed as localhost
	serverB := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bounce" {
			http.Redirect(w, r, addrA+"/back", http.StatusFound)
			return
		}
		record(w, r)
	}))
	defer serverB.Close()
	serverA := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/same":
			http.Redirect(w, r, addrA+"/stayed", http.StatusFound)
		case "/cross":
			http.Redirect(w, r, serverB.URL+"/landed", http.StatusFound)
		case "/chain":
			http.Redirect(w, r, serverB.URL+"/bounce", http.StatusFound)
		default:
			record(w, r)
		}
	}))
	defer serverA.Close()
	u, err := url.Parse(serverA.URL)
	if err != nil {
		t.Fatal(err)
	}
	addrA = "http://localhost:" + u.Port()

	client := &http.Client{Transport: &APIKey{Key: key}}
	tests := []struct {
		path, final, want string
	}{
		// A redirect within the domain keeps the key.
		{"/same", "/stayed", key},
		// A redirect to another domain gets no key.
		{"/cross", "/landed", ""},
		// A chain that left the domain gets no key even where it comes back.
		{"/chain", "/back", ""},
	}
	for _, tt := range tests {
		resp, err := client.Get(addrA + tt.path)
		if err != nil {
			t.Fatalf("GET %s: %v", tt.path, err)
		}
		resp.Body.Close()
		if k := received(tt.final); k != tt.want {
			t.Errorf("GET %s: the redirect target received key %q, want %q", tt.path, k, tt.want)
		}
	}
}

func TestRedirectedCrossDomain(t *testing.T) {
	// chain builds the request an http.Client would hand to a RoundTripper
	// after following redirects through the given URLs, in order.
	chain := func(urls ...string) *http.Request {
		var req *http.Request
		for _, s := range urls {
			u, err := url.Parse(s)
			if err != nil {
				t.Fatal(err)
			}
			next := &http.Request{URL: u, Header: http.Header{}}
			if req != nil {
				next.Response = &http.Response{Request: req}
			}
			req = next
		}
		return req
	}
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{"not a redirect", chain("https://example.com/a"), false},
		{"same host", chain("https://example.com/a", "https://example.com/b"), false},
		{"subdomain", chain("https://example.com/a", "https://api.example.com/b"), false},
		{"parent domain", chain("https://api.example.com/a", "https://example.com/b"), true},
		{"other host", chain("https://example.com/a", "https://evil.example/b"), true},
		{"suffix that is not a subdomain", chain("https://example.com/a", "https://notexample.com/b"), true},
		{"left and came back", chain("https://example.com/a", "https://evil.example/b", "https://example.com/c"), true},
	}
	for _, tt := range tests {
		if got := redirectedCrossDomain(tt.req); got != tt.want {
			t.Errorf("%s: redirectedCrossDomain = %v, want %v", tt.name, got, tt.want)
		}
	}
}
