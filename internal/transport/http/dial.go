// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"context"
	"errors"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	"golang.org/x/oauth2"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/internal"
	"google.golang.org/api/internal/transport/http/propagation"
)

// NewTransport creates an http.RoundTripper for use communicating with a Google
// Cloud service, configured with the given ClientOptions. Its RoundTrip method delegates to base.
func NewTransport(ctx context.Context, base http.RoundTripper, settings *internal.DialSettings) (http.RoundTripper, error) {
	paramTransport := &parameterTransport{
		base:          base,
		userAgent:     settings.UserAgent,
		quotaProject:  settings.QuotaProject,
		requestReason: settings.RequestReason,
	}
	var trans http.RoundTripper = paramTransport
	trans = addOCTransport(trans, settings)
	switch {
	case settings.NoAuth:
		// Do nothing.
	case settings.APIKey != "":
		trans = &transport.APIKey{
			Transport: trans,
			Key:       settings.APIKey,
		}
	default:
		creds, err := internal.Creds(ctx, settings)
		if err != nil {
			return nil, err
		}
		if paramTransport.quotaProject == "" {
			paramTransport.quotaProject = internal.QuotaProjectFromCreds(creds)
		}

		ts := creds.TokenSource
		if settings.TokenSource != nil {
			ts = settings.TokenSource
		}
		trans = &oauth2.Transport{
			Base:   trans,
			Source: ts,
		}
	}
	return trans, nil
}

type parameterTransport struct {
	userAgent     string
	quotaProject  string
	requestReason string

	base http.RoundTripper
}

func (t *parameterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.base
	if rt == nil {
		return nil, errors.New("transport: no Transport specified")
	}
	newReq := *req
	newReq.Header = make(http.Header)
	for k, vv := range req.Header {
		newReq.Header[k] = vv
	}
	if t.userAgent != "" {
		// TODO(cbro): append to existing User-Agent header?
		newReq.Header.Set("User-Agent", t.userAgent)
	}

	// Attach system parameters into the header
	if t.quotaProject != "" {
		newReq.Header.Set("X-Goog-User-Project", t.quotaProject)
	}
	if t.requestReason != "" {
		newReq.Header.Set("X-Goog-Request-Reason", t.requestReason)
	}

	return rt.RoundTrip(&newReq)
}

func addOCTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
	if settings.TelemetryDisabled {
		return trans
	}
	return &ochttp.Transport{
		Base:        trans,
		Propagation: &propagation.HTTPFormat{},
	}
}
