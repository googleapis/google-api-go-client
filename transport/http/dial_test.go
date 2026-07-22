// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/internal"
)

func TestNewClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	client, endpoint, err := NewClient(context.Background())

	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}
	if client == nil {
		t.Fatalf("client is nil")
	}
	if endpoint != "" {
		t.Errorf("got: %s, want: ''", endpoint)
	}
	if got, want := fmt.Sprintf("%T", client.Transport), "*httptransport.authTransport"; got != want {
		t.Fatalf("got %s, want: %s", got, want)
	}
}

func TestHTTP2ReadIdleTimeout(t *testing.T) {
	tests := []struct {
		name     string
		settings *internal.DialSettings
		want     time.Duration
	}{
		{
			name:     "nil settings uses default",
			settings: nil,
			want:     defaultHTTP2ReadIdleTimeout,
		},
		{
			name:     "unset override uses default",
			settings: &internal.DialSettings{},
			want:     defaultHTTP2ReadIdleTimeout,
		},
		{
			name:     "zero override uses default",
			settings: &internal.DialSettings{HTTP2ReadIdleTimeout: 0},
			want:     defaultHTTP2ReadIdleTimeout,
		},
		{
			name:     "negative override uses default",
			settings: &internal.DialSettings{HTTP2ReadIdleTimeout: -1 * time.Second},
			want:     defaultHTTP2ReadIdleTimeout,
		},
		{
			name:     "custom override is honored",
			settings: &internal.DialSettings{HTTP2ReadIdleTimeout: 10 * time.Second},
			want:     10 * time.Second,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := http2ReadIdleTimeout(tc.settings); got != tc.want {
				t.Errorf("http2ReadIdleTimeout() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestConfigureHTTP2AppliesTimeout verifies that configureHTTP2 wires the
// resolved ReadIdleTimeout onto the underlying HTTP/2 transport.
func TestConfigureHTTP2AppliesTimeout(t *testing.T) {
	tests := []struct {
		name     string
		settings *internal.DialSettings
		want     time.Duration
	}{
		{
			name:     "default when unset",
			settings: &internal.DialSettings{},
			want:     defaultHTTP2ReadIdleTimeout,
		},
		{
			name:     "override is applied",
			settings: &internal.DialSettings{HTTP2ReadIdleTimeout: 7 * time.Second},
			want:     7 * time.Second,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// A fresh *http.Transport per case: http2.ConfigureTransports
			// errors if the same transport is configured twice.
			trans := fallbackBaseTransport()
			http2Trans := configureHTTP2(trans, tc.settings)
			if http2Trans == nil {
				t.Fatal("configureHTTP2 returned nil http2 transport")
			}
			if got := http2Trans.ReadIdleTimeout; got != tc.want {
				t.Errorf("ReadIdleTimeout = %v, want %v", got, tc.want)
			}
			if trans.TLSNextProto["h2"] == nil {
				t.Error("configureHTTP2 did not register the h2 next-proto handler")
			}
		})
	}
}
