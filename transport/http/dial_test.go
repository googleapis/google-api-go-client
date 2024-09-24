// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"cloud.google.com/go/auth/httptransport"
	"google.golang.org/api/option/internaloption"
)

func TestNewClient(t *testing.T) {
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

func TestNewClient_NewAuthUniverseDomainEnvVar(t *testing.T) {
	universeDomain := "example.com"
	t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", universeDomain)

	// Replace package var in order to assert DialContext args.
	oldNewClient := newClient
	newClient = func(opts *httptransport.Options) (*http.Client, error) {
		if got, want := opts.UniverseDomain, universeDomain; got != want {
			t.Fatalf("got %s, want: %s", got, want)
		}
		return nil, nil
	}
	defer func() {
		newClient = oldNewClient
	}()

	_, _, err := NewClient(context.Background(), internaloption.EnableNewAuthLibrary())
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}
}
