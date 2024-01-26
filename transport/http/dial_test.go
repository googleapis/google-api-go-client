// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"go.opencensus.io/plugin/ochttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
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
	if got, want := fmt.Sprintf("%T", client.Transport), "*oauth2.Transport"; got != want {
		t.Fatalf("got %s, want: %s", got, want)
	}
	t1 := client.Transport.(*oauth2.Transport)
	if got, want := fmt.Sprintf("%T", t1.Base), "*ochttp.Transport"; got != want {
		t.Fatalf("got %s, want: %s", got, want)
	}
	t2 := t1.Base.(*ochttp.Transport)
	if got, want := fmt.Sprintf("%T", t2.Base), "*otelhttp.Transport"; got != want {
		t.Fatalf("got %s, want: %s", got, want)
	}
}

func TestNewClient_MismatchedUniverseChecks(t *testing.T) {

	rootTokenScope := "https://www.googleapis.com/auth/cloud-platform"
	otherUniverse := "example.com"
	defaultUniverse := "googleapis.com"
	fakeCreds := `
	{"type": "service_account",
     "project_id": "some-project",
     "universe_domain": "UNIVERSE"}`

	// utility function to make a fake credential quickly
	makeFakeCredF := func(universe string) option.ClientOption {
		data := []byte(strings.ReplaceAll(fakeCreds, "UNIVERSE", universe))
		creds, _ := google.CredentialsFromJSON(context.Background(), data, rootTokenScope)
		return option.WithCredentials(creds)
	}

	testCases := []struct {
		description string
		opts        []option.ClientOption
		wantErr     bool
	}{
		{
			description: "default creds and no universe",
			opts: []option.ClientOption{
				option.WithCredentials(&google.Credentials{}),
			},
			wantErr: false,
		},
		{
			description: "default creds and default universe",
			opts: []option.ClientOption{
				option.WithCredentials(&google.Credentials{}),
				option.WithUniverseDomain(defaultUniverse),
			},
			wantErr: false,
		},
		{
			description: "default creds and mismatched universe",
			opts: []option.ClientOption{
				option.WithCredentials(&google.Credentials{}),
				option.WithUniverseDomain(otherUniverse),
			},
			wantErr: true,
		},
		{
			description: "foreign universe creds and default universe",
			opts: []option.ClientOption{
				makeFakeCredF(otherUniverse),
				option.WithUniverseDomain(defaultUniverse),
			},
			wantErr: true,
		},
		{
			description: "foreign universe creds and foreign universe",
			opts: []option.ClientOption{
				makeFakeCredF(otherUniverse),
				option.WithUniverseDomain(otherUniverse),
			},
			wantErr: false,
		},
		{
			description: "tokensource + mismatched universe",
			opts: []option.ClientOption{
				option.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{})),
				option.WithUniverseDomain(otherUniverse),
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		opts := []option.ClientOption{
			option.WithScopes(rootTokenScope),
		}
		opts = append(opts, tc.opts...)
		_, _, gotErr := NewClient(context.Background(), opts...)
		if tc.wantErr && gotErr == nil {
			t.Errorf("%q: wanted error, got none", tc.description)
		}
		if !tc.wantErr && gotErr != nil {
			t.Errorf("%q: wanted success, got err: %v", tc.description, gotErr)
		}
	}
}
