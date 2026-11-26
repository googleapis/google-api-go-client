// Copyright 2012 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transport contains HTTP transports used to make
// authenticated API requests.
//
// This package is DEPRECATED. Users should instead use,
//
//	service, err := NewService(..., option.WithAPIKey(...))
package transport

import (
	"errors"
	"net/http"
	"strings"
)

// APIKey is an HTTP Transport which wraps an underlying transport and
// appends an API Key "key" parameter to the URL of outgoing requests.
//
// A request that an [http.Client] created to follow a redirect away from
// the domain of the original request (or a subdomain of it) is sent without
// the key, so that a redirect cannot hand the key to a host the caller
// never configured. This is the rule net/http applies to the Authorization
// header on such a redirect.
//
// Deprecated: please use NewService(..., option.WithAPIKey(...)) instead.
type APIKey struct {
	// Key is the API Key to set on requests.
	Key string

	// Transport is the underlying HTTP transport.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper
}

func (t *APIKey) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
		if rt == nil {
			return nil, errors.New("googleapi/transport: no Transport specified or available")
		}
	}
	if redirectedCrossDomain(req) {
		return rt.RoundTrip(req)
	}
	// Per the RoundTripper contract the caller's request stays untouched,
	// including the URL it points to: the key is added to a copy, so it never
	// reaches an error message that prints the caller's URL.
	newReq := *req
	if req.URL != nil {
		newURL := *req.URL
		newReq.URL = &newURL
	}
	args := newReq.URL.Query()
	args.Set("key", t.Key)
	newReq.URL.RawQuery = args.Encode()
	return rt.RoundTrip(&newReq)
}

// redirectedCrossDomain reports whether req was created by an http.Client
// to follow a redirect chain that has left the domain of the request that
// started it. The client records the redirect response in req.Response, so
// following that chain back reaches the original request. Once any hop has
// left the original domain, every later hop counts as cross-domain as well,
// even one that returns to it, which is how net/http treats the
// Authorization header.
//
// Hostnames are compared in lower case as they appear in the URLs. net/http
// also converts an internationalized hostname to its ASCII form first; a
// chain that mixes the two forms of one name is treated as cross-domain
// here, which withholds the key rather than sending it.
func redirectedCrossDomain(req *http.Request) bool {
	if req.Response == nil {
		return false
	}
	first := req
	for first.Response != nil && first.Response.Request != nil {
		first = first.Response.Request
	}
	domain := strings.ToLower(first.URL.Hostname())
	for r := req; r != first; r = r.Response.Request {
		if !isDomainOrSubdomain(strings.ToLower(r.URL.Hostname()), domain) {
			return true
		}
	}
	return false
}

// isDomainOrSubdomain reports whether sub is a subdomain (or exact match) of
// the parent domain. It is the rule net/http applies when deciding whether a
// redirect may keep the Authorization header, copied from net/http/client.go.
// Both domains must already be in canonical form.
func isDomainOrSubdomain(sub, parent string) bool {
	if sub == parent {
		return true
	}
	// If sub contains a :, it's probably an IPv6 address (and is definitely not a hostname).
	// Don't check the suffix in this case, to avoid matching the contents of a IPv6 zone.
	// For example, "::1%.www.example.com" is not a subdomain of "www.example.com".
	if strings.ContainsAny(sub, ":%") {
		return false
	}
	// If sub is "foo.example.com" and parent is "example.com",
	// that means sub must end in "."+parent.
	// Do it without allocating.
	if !strings.HasSuffix(sub, parent) {
		return false
	}
	return sub[len(sub)-len(parent)-1] == '.'
}
