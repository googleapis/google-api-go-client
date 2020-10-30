// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// Check that user optioned grpc.WithDialer option overwrites App Engine dialer
func TestGRPCHook(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	expected := false

	appengineDialerHook = (func(ctx context.Context) grpc.DialOption {
		return grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			t.Error("did not expect a call to appengine dialer, got one")
			cancel()
			return nil, errors.New("not expected")
		})
	})
	defer func() {
		appengineDialerHook = nil
	}()

	expectedDialer := grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
		expected = true
		cancel()
		return nil, errors.New("expected")
	})

	conn, err := Dial(ctx,
		option.WithTokenSource(oauth2.StaticTokenSource(nil)), // No creds.
		option.WithGRPCDialOption(expectedDialer),
		option.WithGRPCDialOption(grpc.WithBlock()))
	if err != context.Canceled {
		t.Errorf("got %v, want %v", err, context.Canceled)
	}
	if conn != nil {
		conn.Close()
		t.Error("got valid conn, want nil")
	}
	if !expected {
		t.Error("expected a call to expected dialer, didn't get one")
	}
}

func TestCheckDirectPathEndPoint(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		endpoint string
		want     bool
	}{
		{
			name:     "empty endpoint are disallowed",
			endpoint: "",
			want:     false,
		},
		{
			name:     "dns schemes are allowed",
			endpoint: "dns:///foo",
			want:     true,
		},
		{
			name:     "host without no prefix are allowed",
			endpoint: "foo",
			want:     true,
		},
		{
			name:     "host with port are allowed",
			endpoint: "foo:1234",
			want:     true,
		},
		{
			name:     "non-dns schemes are disallowed",
			endpoint: "https://foo",
			want:     false,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			if got := checkDirectPathEndPoint(testcase.endpoint); got != testcase.want {
				t.Fatalf("got %v, want %v", got, testcase.want)
			}
		})
	}
}
