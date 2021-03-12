// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package idtoken

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

const (
	envCredentialFile = "GCLOUD_TESTS_GOLANG_KEY"

	aud = "http://example.com"
)

func TestNewTokenSource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	os.Setenv(envCredentialFile, "/Users/codyoss/creds/codyoss-workspace-2b1fb7cd40c0.json")
	ts, err := idtoken.NewTokenSource(context.Background(), "http://example.com", option.WithCredentialsFile(os.Getenv(envCredentialFile)))
	if err != nil {
		t.Fatalf("unable to create TokenSource: %v", err)
	}
	tok, err := ts.Token()
	if err != nil {
		t.Fatalf("unable to retrieve Token: %v", err)
	}
	req := &http.Request{Header: make(http.Header)}
	tok.SetAuthHeader(req)
	if !strings.HasPrefix(req.Header.Get("Authorization"), "Bearer ") {
		t.Fatalf("token should sign requests with Bearer Authorization header")
	}
	validTok, err := idtoken.Validate(context.Background(), tok.AccessToken, aud)
	if err != nil {
		t.Fatalf("token validation failed: %v", err)
	}
	if validTok.Audience != aud {
		t.Fatalf("got %q, want %q", validTok.Audience, aud)
	}
}

func TestNewClient_WithCredentialFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	os.Setenv(envCredentialFile, "/Users/codyoss/creds/codyoss-workspace-2b1fb7cd40c0.json")
	client, err := idtoken.NewClient(context.Background(), aud, option.WithCredentialsFile(os.Getenv(envCredentialFile)))
	if err != nil {
		t.Fatalf("unable to create Client: %v", err)
	}
	tok, err := client.Transport.(*oauth2.Transport).Source.Token()
	if err != nil {
		t.Fatalf("unable to retrieve Token: %v", err)
	}
	validTok, err := idtoken.Validate(context.Background(), tok.AccessToken, aud)
	if err != nil {
		t.Fatalf("token validation failed: %v", err)
	}
	if validTok.Audience != aud {
		t.Fatalf("got %q, want %q", validTok.Audience, aud)
	}
}

func TestNewClient_WithCredentialJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.Background()
	creds, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		t.Fatalf("unable to find default creds: %v", err)
	}
	client, err := idtoken.NewClient(ctx, aud, option.WithCredentialsJSON(creds.JSON))
	if err != nil {
		t.Fatalf("unable to create Client: %v", err)
	}
	tok, err := client.Transport.(*oauth2.Transport).Source.Token()
	if err != nil {
		t.Fatalf("unable to retrieve Token: %v", err)
	}
	validTok, err := idtoken.Validate(context.Background(), tok.AccessToken, aud)
	if err != nil {
		t.Fatalf("token validation failed: %v", err)
	}
	if validTok.Audience != aud {
		t.Fatalf("got %q, want %q", validTok.Audience, aud)
	}
}
