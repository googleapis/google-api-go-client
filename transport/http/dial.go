// Copyright 2015 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package http supports network connections to HTTP servers.
// This package is not intended for use by end developers. Use the
// google.golang.org/api/option package to configure API clients.
package http

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"go.opencensus.io/plugin/ochttp"
	"golang.org/x/oauth2"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/http/internal/propagation"
)

// NewClient returns an HTTP client for use communicating with a Google cloud
// service, configured with the given ClientOptions. It also returns the endpoint
// for the service as specified in the options.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*http.Client, string, error) {
	settings, err := newSettings(opts)
	if err != nil {
		return nil, "", err
	}
	endpoint, err := getEndpoint(settings)
	if err != nil {
		return nil, "", err
	}
	// TODO(cbro): consider injecting the User-Agent even if an explicit HTTP client is provided?
	if settings.HTTPClient != nil {
		return settings.HTTPClient, endpoint, nil
	}
	trans, err := newTransport(ctx, defaultBaseTransport(ctx), settings)
	if err != nil {
		return nil, "", err
	}
	return &http.Client{Transport: trans}, endpoint, nil
}

// NewTransport creates an http.RoundTripper for use communicating with a Google
// cloud service, configured with the given ClientOptions. Its RoundTrip method delegates to base.
func NewTransport(ctx context.Context, base http.RoundTripper, opts ...option.ClientOption) (http.RoundTripper, error) {
	settings, err := newSettings(opts)
	if err != nil {
		return nil, err
	}
	if settings.HTTPClient != nil {
		return nil, errors.New("transport/http: WithHTTPClient passed to NewTransport")
	}
	return newTransport(ctx, base, settings)
}

func newTransport(ctx context.Context, base http.RoundTripper, settings *internal.DialSettings) (http.RoundTripper, error) {
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
		trans = &oauth2.Transport{
			Base:   trans,
			Source: creds.TokenSource,
		}
	}
	return trans, nil
}

func newSettings(opts []option.ClientOption) (*internal.DialSettings, error) {
	var o internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&o)
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if o.GRPCConn != nil {
		return nil, errors.New("unsupported gRPC connection specified")
	}
	return &o, nil
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

// Set at init time by dial_appengine.go. If nil, we're not on App Engine.
var appengineUrlfetchHook func(context.Context) http.RoundTripper

// defaultBaseTransport returns the base HTTP transport.
// On App Engine, this is urlfetch.Transport, otherwise it's http.DefaultTransport.
func defaultBaseTransport(ctx context.Context) http.RoundTripper {
	if appengineUrlfetchHook != nil {
		return appengineUrlfetchHook(ctx)
	}
	return http.DefaultTransport
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

// getEndpoint gets the endpoint for the service.
//
// If the user-provided endpoint is an address (host:port) rather than full base
// URL (https://...), then the user-provided address is merged into the default
// endpoint.
//
// For example, (WithEndpoint("myhost:8000"), WithDefaultEndpoint("https://foo.com/bar/baz")) will return "https://myhost:8080/bar/baz"
func getEndpoint(settings *internal.DialSettings) (string, error) {
	if settings.Endpoint == "" {
		return settings.DefaultEndpoint, nil
	}
	if strings.Contains(settings.Endpoint, "://") {
		// User passed in a full URL path, use it verbatim.
		return settings.Endpoint, nil
	}
	if settings.DefaultEndpoint == "" {
		return "", errors.New("WithEndpoint requires a full URL path")
	}

	// Assume user-provided endpoint is host[:port], merge it with the default endpoint.
	return mergeEndpoints(settings.DefaultEndpoint, settings.Endpoint)
}

func mergeEndpoints(base, newHost string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	u.Host = newHost
	return u.String(), nil
}
