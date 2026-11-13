// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/api/option"
)

func TestNewClient_WithHTTPClient_IgnoresClientCertConfig(t *testing.T) {
	tempDir := t.TempDir()
	badConfig := filepath.Join(tempDir, "bad_cert.json")
	if err := os.WriteFile(badConfig, []byte("{"), 0600); err != nil {
		t.Fatalf("failed to write bad cert config: %v", err)
	}
	t.Setenv("GOOGLE_API_CERTIFICATE_CONFIG", badConfig)
	t.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "true")

	customClient := &http.Client{}
	endpoint := "https://www.googleapis.com/youtube/v3/"
	client, gotEndpoint, err := NewClient(context.Background(),
		option.WithHTTPClient(customClient),
		option.WithEndpoint(endpoint),
	)
	if err != nil {
		t.Fatalf("unexpected error with injected HTTPClient: %v", err)
	}
	if client != customClient {
		t.Errorf("got client %p, want %p", client, customClient)
	}
	if gotEndpoint != endpoint {
		t.Errorf("got endpoint %q, want %q", gotEndpoint, endpoint)
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
