// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package impersonate is used to impersonate Google Credentials.
package impersonate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// Config for generating impersonated credentials.
type Config struct {
	// Target is the service account to impersonate. Required.
	Target string
	// Scopes the impersonated credential should have. Required for access tokens.
	Scopes []string
	// Audience for the token. Required for ID tokens.
	Audience string
	// IncludeEmail in the generated ID token. Optional.
	IncludeEmail bool
	// Delegates are the service accounts in a delegation chain. Each service
	// account must be granted roles/iam.serviceAccountTokenCreator on the next
	// service account in the chain. Optional.
	Delegates []string
}

// TokenSource returns an impersonated TokenSource configured with the provided
// config using ts as the base credential provider for making requests.
func TokenSource(ctx context.Context, ts oauth2.TokenSource, config *Config) (oauth2.TokenSource, error) {
	if len(config.Scopes) > 0 {
		return accessTokenSource(ctx, ts, config)
	} else if config.Audience != "" {
		return idTokenSource(ctx, ts, config)
	}

	return nil, fmt.Errorf("impersonate: scopes or audience must be provided")
}

func accessTokenSource(ctx context.Context, ts oauth2.TokenSource, config *Config) (oauth2.TokenSource, error) {
	its := impersonatedAccessTokenSource{
		ctx:  ctx,
		ts:   ts,
		name: formatIAMServiceAccountName(config.Target),
		// Default to the longest acceptable value of one hour as the token will
		// be refreshed automatically.
		lifetime: "3600s",
	}

	its.delegates = make([]string, len(config.Delegates))
	for i, v := range config.Delegates {
		its.delegates[i] = formatIAMServiceAccountName(v)
	}
	its.scopes = make([]string, len(config.Scopes))
	copy(its.scopes, config.Scopes)

	return oauth2.ReuseTokenSource(nil, its), nil
}

func idTokenSource(ctx context.Context, ts oauth2.TokenSource, config *Config) (oauth2.TokenSource, error) {
	its := impersonatedIdTokenSource{
		ctx:          ctx,
		ts:           ts,
		name:         formatIAMServiceAccountName(config.Target),
		audience:     config.Audience,
		includeEmail: config.IncludeEmail,
	}

	its.delegates = make([]string, len(config.Delegates))
	for i, v := range config.Delegates {
		its.delegates[i] = formatIAMServiceAccountName(v)
	}

	return oauth2.ReuseTokenSource(nil, its), nil
}

func formatIAMServiceAccountName(name string) string {
	return fmt.Sprintf("projects/-/serviceAccounts/%s", name)
}

type generateAccessTokenReq struct {
	Delegates []string `json:"delegates,omitempty"`
	Lifetime  string   `json:"lifetime,omitempty"`
	Scope     []string `json:"scope,omitempty"`
}

type generateAccessTokenResp struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  string `json:"expireTime"`
}

type impersonatedAccessTokenSource struct {
	ctx context.Context
	ts  oauth2.TokenSource

	name      string
	lifetime  string
	scopes    []string
	delegates []string
}

// Token returns an impersonated Token.
func (i impersonatedAccessTokenSource) Token() (*oauth2.Token, error) {
	req := generateAccessTokenReq{
		Delegates: i.delegates,
		Lifetime:  i.lifetime,
		Scope:     i.scopes,
	}
	resp := generateAccessTokenResp{}

	hc := oauth2.NewClient(i.ctx, i.ts)
	url := fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateAccessToken", i.name)
	if err := getToken(i.ctx, hc, url, &req, &resp); err != nil {
		return nil, err
	}

	expiry, err := time.Parse(time.RFC3339, resp.ExpireTime)
	if err != nil {
		return nil, fmt.Errorf("impersonate: unable to parse expiry: %v", err)
	}

	return &oauth2.Token{
		AccessToken: resp.AccessToken,
		Expiry:      expiry,
	}, nil
}

type impersonatedIdTokenSource struct {
	ctx context.Context
	ts  oauth2.TokenSource

	name         string
	audience     string
	delegates    []string
	includeEmail bool
}

type generateIdTokenReq struct {
	Audience     string   `json:"audience,omitempty"`
	Delegates    []string `json:"delegates,omitempty"`
	IncludeEmail bool     `json:"includeEmail,omitempty"`
}

type generateIdTokenResp struct {
	Token string `json:"token"`
}

// Token returns an impersonated Token.
func (i impersonatedIdTokenSource) Token() (*oauth2.Token, error) {
	req := generateIdTokenReq{
		Audience:     i.audience,
		Delegates:    i.delegates,
		IncludeEmail: i.includeEmail,
	}
	resp := generateIdTokenResp{}

	hc := oauth2.NewClient(i.ctx, i.ts)
	url := fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateIdToken", i.name)
	if err := getToken(i.ctx, hc, url, &req, &resp); err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: resp.Token,
		// ID tokens are valid for one hour, leave a little buffer
		Expiry: time.Now().Add(55 * time.Minute),
	}, nil
}

func getToken(ctx context.Context, hc *http.Client, url string, in interface{}, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("impersonate: unable to marshal request: %v", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("impersonate: unable to create request: %v", err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		return fmt.Errorf("impersonate: unable to generate token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return fmt.Errorf("impersonate: unable to read body: %v", err)
	}
	if c := resp.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("impersonate: status code %d: %s", c, body)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("impersonate: unable to parse response: %v", err)
	}
	return nil
}
