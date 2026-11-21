// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
)

func TestNewClientWithHTTPClient(t *testing.T) {
	// The default certificate source is cached process-wide. Run in a fresh
	// process so another test cannot hide an attempt to initialize it.
	const childEnv = "GOOGLE_API_TEST_HTTP_CLIENT_CERT_CHILD"
	if os.Getenv(childEnv) != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestNewClientWithHTTPClient$", "-test.v")
		cmd.Env = append(os.Environ(), childEnv+"=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("certificate isolation test: %v\n%s", err, out)
		}
		return
	}
	config := filepath.Join(t.TempDir(), "certificate.json")
	if err := os.WriteFile(config, []byte("{"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GOOGLE_API_CERTIFICATE_CONFIG", config)
	t.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "true")
	t.Setenv("EXPERIMENTAL_GOOGLE_API_USE_S2A", "true")
	client := &http.Client{}
	for _, tc := range []struct {
		name, endpoint, universe, mtls, want string
		wantErr                              bool
	}{
		{name: "explicit URL", endpoint: "https://custom.example/v4/", mtls: "auto", want: "https://custom.example/v4/"},
		{name: "default", mtls: "auto", want: "https://test.googleapis.com/v4/"},
		{name: "host override", endpoint: "localhost:8080", mtls: "auto", want: "https://localhost:8080/v4/"},
		{name: "universe", universe: "example.com", mtls: "auto", want: "https://test.example.com/v4/"},
		{name: "explicit mTLS mode", mtls: "always", want: "https://test.mtls.googleapis.com/v4/"},
		{name: "unsupported mTLS universe", universe: "example.com", mtls: "always", wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("GOOGLE_API_USE_MTLS_ENDPOINT", tc.mtls)
			opts := []option.ClientOption{
				option.WithHTTPClient(client),
				internaloption.WithDefaultEndpoint("https://test.googleapis.com/v4/"),
				internaloption.WithDefaultEndpointTemplate("https://test.UNIVERSE_DOMAIN/v4/"),
				internaloption.WithDefaultMTLSEndpoint("https://test.mtls.googleapis.com/v4/"),
			}
			if tc.endpoint != "" {
				opts = append(opts, option.WithEndpoint(tc.endpoint))
			}
			if tc.universe != "" {
				opts = append(opts, option.WithUniverseDomain(tc.universe))
			}
			got, endpoint, err := NewClient(context.Background(), opts...)
			if tc.wantErr {
				if err == nil || err.Error() != "mTLS is not supported in any universe other than googleapis.com" {
					t.Fatalf("got error %v, want unsupported mTLS universe", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != client || endpoint != tc.want {
				t.Fatalf("got client %p, endpoint %q; want client %p, endpoint %q", got, endpoint, client, tc.want)
			}
		})
	}
	// SDK-managed transports must still report malformed certificate settings.
	t.Setenv("GOOGLE_API_USE_MTLS_ENDPOINT", "auto")
	if _, _, err := NewClient(context.Background(), option.WithoutAuthentication()); err == nil || err.Error() != "unexpected end of JSON input" {
		t.Fatalf("SDK-managed transport error = %v, want malformed certificate error", err)
	}
}

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
