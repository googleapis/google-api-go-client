// Copyright 2017 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal/impersonate"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	creds, err := baseCreds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if ds.ImpersonationConfig != nil {
		return impersonateCredentials(ctx, creds, ds)
	}
	return creds, nil
}

func baseCreds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials based on the input.
//
// - A standard OAuth 2.0 flow will be executed if at least one of the
//   following conditions is met:
//   (1) the scope is non-empty and the scope for self-signed JWT flow is
//       disabled.
//   (2) Service Account Impersontation
//
// - Otherwise, executes a self-signed JWT flow (google.aip.dev/auth/4111)
func credentialsFromJSON(ctx context.Context, data []byte, ds *DialSettings) (*google.Credentials, error) {
	cred, err := google.CredentialsFromJSON(ctx, data, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if isOAuthFlow(data, ds) {
		// Standard OAuth 2.0 Flow
		return cred, nil
	}

	isJWTFlow, err := isSelfSignedJWTFlow(cred.JSON)
	if err != nil {
		return nil, err
	}

	if isJWTFlow {
		ts, err := selfSignedJWTTokenSource(data, ds.GetAudience(), ds.GetScopes())
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}
	return cred, err
}

func isOAuthFlow(data []byte, ds *DialSettings) bool {
	// Standard OAuth 2.0 Flow
	return len(data) == 0 ||
		(len(ds.GetScopes()) > 0 && !ds.EnableScopeForJWT) ||
		ds.ImpersonationConfig != nil
}

func isSelfSignedJWTFlow(data []byte) (bool, error) {
	// Check if JSON is a service account and if so create a self-signed JWT.
	var f struct {
		Type string `json:"type"`
		// The rest JSON fields are omitted because they are not used.
	}
	if err := json.Unmarshal(data, &f); err != nil {
		return false, err
	}
	return f.Type == serviceAccountKey, nil
}

func selfSignedJWTTokenSource(data []byte, audience string, scopes []string) (oauth2.TokenSource, error) {
	if len(scopes) > 0 {
		// Scopes are preferred in self-signed JWT
		return google.JWTAccessTokenSourceWithScope(data, scopes...)
	} else if audience != "" {
		// Fallback to audience if scope is not provided
		return google.JWTAccessTokenSourceFromJSON(data, audience)
	} else {
		return nil, errors.New("neither scopes or audience are provided for the self-signed JWT")
	}
}

// QuotaProjectFromCreds returns the quota project from the JSON blob in the provided credentials.
//
// NOTE(cbro): consider promoting this to a field on google.Credentials.
func QuotaProjectFromCreds(cred *google.Credentials) string {
	var v struct {
		QuotaProject string `json:"quota_project_id"`
	}
	if err := json.Unmarshal(cred.JSON, &v); err != nil {
		return ""
	}
	return v.QuotaProject
}

func impersonateCredentials(ctx context.Context, creds *google.Credentials, ds *DialSettings) (*google.Credentials, error) {
	if len(ds.ImpersonationConfig.Scopes) == 0 {
		ds.ImpersonationConfig.Scopes = ds.GetScopes()
	}
	ts, err := impersonate.TokenSource(ctx, creds.TokenSource, ds.ImpersonationConfig)
	if err != nil {
		return nil, err
	}
	return &google.Credentials{
		TokenSource: ts,
		ProjectID:   creds.ProjectID,
	}, nil
}
