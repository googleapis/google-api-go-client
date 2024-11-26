// Copyright 2015 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transport

import (
	"context"
	"net/http"

	"cloud.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"

	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	htransport "google.golang.org/api/transport/http"
)

// NewHTTPClient returns an HTTP client for use communicating with a Google cloud
// service, configured with the given ClientOptions. It also returns the endpoint
// for the service as specified in the options.
func NewHTTPClient(ctx context.Context, opts ...option.ClientOption) (*http.Client, string, error) {
	return htransport.NewClient(ctx, opts...)
}

// DialGRPC returns a GRPC connection for use communicating with a Google cloud
// service, configured with the given ClientOptions.
func DialGRPC(ctx context.Context, opts ...option.ClientOption) (*grpc.ClientConn, error) {
	return gtransport.Dial(ctx, opts...)
}

// DialGRPCInsecure returns an insecure GRPC connection for use communicating
// with fake or mock Google cloud service implementations, such as emulators.
// The connection is configured with the given ClientOptions.
func DialGRPCInsecure(ctx context.Context, opts ...option.ClientOption) (*grpc.ClientConn, error) {
	return gtransport.DialInsecure(ctx, opts...)
}

// Creds constructs a google.Credentials from the information in the options,
// or obtains the default credentials in the same way as google.FindDefaultCredentials.
func Creds(ctx context.Context, opts ...option.ClientOption) (*google.Credentials, error) {
	var ds internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&ds)
	}
	return internal.Creds(ctx, &ds)
}

// AuthCreds returns [cloud.google.com/go/auth.Credentials] using the following
// options provided via [option.ClientOption], including legacy oauth2/google
// options, in this order:
//
// * [option.WithAuthCredentials]
// * [option/internaloption.WithCredentials] (internal use only)
// * [option.WithCredentials]
// * [option.WithTokenSource]
//
// If there are no applicable credentials options, then it passes the
// following options to [cloud.google.com/go/auth/credentials.DetectDefault] and
// returns the result:
//
// * [option.WithAudiences]
// * [option.WithCredentialsFile]
// * [option.WithCredentialsJSON]
// * [option.WithScopes]
// * [option/internaloption.WithDefaultScopes] (internal use only)
// * [option/internaloption.EnableJwtWithScope] (internal use only)
func AuthCreds(ctx context.Context, opts ...option.ClientOption) (*auth.Credentials, error) {
	var ds internal.DialSettings
	for _, opt := range opts {
		opt.Apply(&ds)
	}
	return internal.AuthCreds(ctx, &ds)
}
