// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"context"
	"net/http"
	"testing"

	"golang.org/x/xerrors"
)

func TestSendRequest(t *testing.T) {
	// Setting Accept-Encoding should give an error immediately.
	req, _ := http.NewRequest("GET", "url", nil)
	req.Header.Set("Accept-Encoding", "")
	_, err := SendRequest(context.Background(), nil, req)
	if err == nil {
		t.Error("got nil, want error")
	}
}

func TestSendRequestWithRetry(t *testing.T) {
	// Setting Accept-Encoding should give an error immediately.
	req, _ := http.NewRequest("GET", "url", nil)
	req.Header.Set("Accept-Encoding", "")
	_, err := SendRequestWithRetry(context.Background(), nil, req, nil)
	if err == nil {
		t.Error("got nil, want error")
	}
}

type brokenRoundTripper struct{}

func (t *brokenRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, xerrors.New("this should not happen")
}

func TestCanceledContextDoesNotPerformRequest(t *testing.T) {
	client := http.Client{
		Transport: &brokenRoundTripper{},
	}
	for i := 0; i < 1000; i++ {
		req, _ := http.NewRequest("GET", "url", nil)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := SendRequestWithRetry(ctx, &client, req, nil)
		if !xerrors.Is(err, context.Canceled) {
			t.Fatalf("got %v, want %v", err, context.Canceled)
		}
	}
}
