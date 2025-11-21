// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idtoken_test

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"

	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

const (
	envCredentialFile = "GOOGLE_APPLICATION_CREDENTIALS"
	// Change this type as needed to match the credentials type of GOOGLE_APPLICATION_CREDENTIALS JSON or ADC credentials JSON.
	credentialsFileType = idtoken.ServiceAccount

	aud = "http://example.com"
)

func TestNewTokenSource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	credsPath := os.Getenv(envCredentialFile)
	if credsPath == "" {
		t.Fatalf("Env var is not set: %s", envCredentialFile)
	}
	ts, err := idtoken.NewTokenSource(context.Background(), "http://example.com", option.WithAuthCredentialsFile(credentialsFileType, credsPath))
	if err != nil {
		t.Fatalf("unable to create Client: %v", err)
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
