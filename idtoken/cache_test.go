// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idtoken

import (
	"net/http"
	"testing"
	"time"
)

func TestCacheHit(t *testing.T) {
	dummyResp := &certResponse{
		Keys: []jwk{
			{
				Kid: "123",
			},
		},
	}
	cache := newCachingClient(nil)
	// Cache should be empty
	cert, ok := cache.get(googleSACertsURL)
	if ok || cert != nil {
		t.Fatal("cache for SA certs should be empty")
	}

	// Add an item, but make it expire now
	cache.set(googleSACertsURL, dummyResp, make(http.Header))
	cert, ok = cache.get(googleSACertsURL)
	if ok || cert != nil {
		t.Fatal("cache for SA certs should be expired")
	}

	// Add an item that expires in 1 seconds
	h := make(http.Header)
	h.Set("age", "0")
	h.Set("cache-control", "public, max-age=1, must-revalidate, no-transform")
	cache.set(googleSACertsURL, dummyResp, h)
	cert, ok = cache.get(googleSACertsURL)
	if !ok || cert == nil || cert.Keys[0].Kid != "123" {
		t.Fatal("cache for SA certs have a resp")
	}
	// Wait
	time.Sleep(2 * time.Second)
	cert, ok = cache.get(googleSACertsURL)
	if ok || cert != nil {
		t.Fatal("cache for SA certs should be expired")
	}
}
