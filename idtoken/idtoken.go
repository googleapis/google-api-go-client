// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idtoken

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
)

// ClientOption is aliased so relevant options are easily found in the docs.

// ClientOption is for configuring a Google API client or transport.
type ClientOption = option.ClientOption

// NewClient creates a HTTP Client that automatically adds an ID token to each
// request via an Authorization header. The token will have have the audience
// provided and be configured with the supplied options. The parameter audience
// may not be empty.
func NewClient(ctx context.Context, audience string, opts ...ClientOption) (*http.Client, error) {
	var ds internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&ds)
	}
	if err := ds.Validate(); err != nil {
		return nil, err
	}
	if ds.NoAuth {
		return nil, fmt.Errorf("idtoken: option.WithoutAuthentication not supported")
	}
	if ds.APIKey != "" {
		return nil, fmt.Errorf("idtoken: option.WithAPIKey not supported")
	}
	if ds.TokenSource != nil {
		return nil, fmt.Errorf("idtoken: option.WithTokenSource not supported")
	}

	ts, err := NewTokenSource(ctx, audience, opts...)
	if err != nil {
		return nil, err
	}
	opts = append(opts, option.WithTokenSource(ts))
	t, err := htransport.NewTransport(ctx, http.DefaultTransport, opts...)
	if err != nil {
		return nil, err
	}
	return &http.Client{Transport: t}, nil
}

// NewTokenSource creates a TokenSource that returns ID tokens with the audience
// provided and configured with the supplied options. The parameter audience may
// not be empty.
func NewTokenSource(ctx context.Context, audience string, opts ...ClientOption) (oauth2.TokenSource, error) {
	if audience == "" {
		return nil, fmt.Errorf("idtoken: must supply a non-empty audience")
	}
	var ds internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&ds)
	}
	if err := ds.Validate(); err != nil {
		return nil, err
	}
	if ds.TokenSource != nil {
		return nil, fmt.Errorf("idtoken: option.WithTokenSource not supported")
	}
	return newTokenSource(ctx, audience, &ds)
}

func newTokenSource(ctx context.Context, audience string, ds *internal.DialSettings) (oauth2.TokenSource, error) {
	creds, err := internal.Creds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if len(creds.JSON) > 0 {
		return tokenSourceFromBytes(ctx, creds.JSON, audience, ds)
	}
	// If internal.Creds did not return a response with JSON fallback to the
	// metadata service as the creds.TokenSource is not an ID token.
	if metadata.OnGCE() {
		return computeTokenSource(audience, ds)
	}
	return nil, fmt.Errorf("idtoken: couldn't find any credentials")
}

func tokenSourceFromBytes(ctx context.Context, data []byte, audience string, ds *internal.DialSettings) (oauth2.TokenSource, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("idtoken: credential provided is 0 bytes")
	}

	var f struct {
		Type         string `json:"type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
	}

	fmt.Println(string(data))

	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}

	switch f.Type {
	case "service_account":
		cfg, err := google.JWTConfigFromJSON(data, ds.Scopes...)
		if err != nil {
			return nil, err
		}

		customClaims := ds.CustomClaims
		if customClaims == nil {
			customClaims = make(map[string]interface{})
		}
		customClaims["target_audience"] = audience

		cfg.PrivateClaims = customClaims
		cfg.UseIDToken = true

		ts := cfg.TokenSource(ctx)
		tok, err := ts.Token()
		if err != nil {
			return nil, err
		}

		return oauth2.ReuseTokenSource(tok, ts), nil
	case "authorized_user":
		refresh := oauth2.Token{
			RefreshToken: f.RefreshToken,
		}

		cfg := &oauth2.Config{
			ClientID:     f.ClientID,
			ClientSecret: f.ClientSecret,
			Endpoint:     google.Endpoint,
			Scopes:       ds.Scopes,
		}

		cfg.TokenSource(ctx, &refresh)

		ts := idTokenSource{cfg.TokenSource(ctx, &refresh)}

		tok, err := ts.Token()
		if err != nil {
			return nil, err
		}

		return oauth2.ReuseTokenSource(tok, ts), nil
	default:
		return nil, fmt.Errorf("idtoken: unsupported credential type %q", f.Type)
	}
}

type idTokenSource struct {
	ts oauth2.TokenSource
}

func (s idTokenSource) Token() (*oauth2.Token, error) {
	tok, err := s.ts.Token()
	if err != nil {
		return nil, err
	}

	rawID := tok.Extra("id_token")
	if rawID == nil {
		return tok, nil
	}

	id, ok := rawID.(string)
	if !ok {
		return tok, nil
	}

	tok.AccessToken = id

	return tok, nil
}

// WithCustomClaims optionally specifies custom private claims for an ID token.
func WithCustomClaims(customClaims map[string]interface{}) ClientOption {
	return withCustomClaims(customClaims)
}

type withCustomClaims map[string]interface{}

func (w withCustomClaims) Apply(o *internal.DialSettings) {
	o.CustomClaims = w
}

// WithCredentialsFile returns a ClientOption that authenticates
// API calls with the given service account or refresh token JSON
// credentials file.
func WithCredentialsFile(filename string) ClientOption {
	return option.WithCredentialsFile(filename)
}

// WithCredentialsJSON returns a ClientOption that authenticates
// API calls with the given service account or refresh token JSON
// credentials.
func WithCredentialsJSON(p []byte) ClientOption {
	return option.WithCredentialsJSON(p)
}

// WithHTTPClient returns a ClientOption that specifies the HTTP client to use
// as the basis of communications. This option may only be used with services
// that support HTTP as their communication transport. When used, the
// WithHTTPClient option takes precedent over all other supplied options.
func WithHTTPClient(client *http.Client) ClientOption {
	return option.WithHTTPClient(client)
}
