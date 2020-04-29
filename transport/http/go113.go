// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.13

package http

import (
	"context"
	"crypto/tls"
	"net/http"

	"google.golang.org/api/transport/cert"
)

// defaultBaseTransport returns the base HTTP transport.
// On App Engine, this is urlfetch.Transport.
// Otherwise, use a default transport, taking most defaults from
// http.DefaultTransport.
// If TLSCertificate is available, set TLSClientConfig as well.
func defaultBaseTransport(ctx context.Context, clientCertSource cert.Source) http.RoundTripper {
	if appengineUrlfetchHook != nil {
		return appengineUrlfetchHook(ctx)
	}

	// Copy http.DefaultTransport except for MaxIdleConnsPerHost setting,
	// which is increased due to reported performance issues under load in the GCS
	// client. Transport.Clone is only available in Go 1.13 and up.
	trans := http.DefaultTransport.(*http.Transport).Clone()
	trans.MaxIdleConnsPerHost = 100

	if clientCertSource != nil {
		trans.TLSClientConfig = &tls.Config{
			GetClientCertificate: clientCertSource,
		}
	}

	return trans
}
