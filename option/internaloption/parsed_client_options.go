// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internaloption contains options used internally by Google client code.
package internaloption

import (
	"crypto/tls"
	"log/slog"
	"net/http"

	"cloud.google.com/go/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// ParsedClientOptions provides read-only access to settings derived from a list of ClientOption.
// It is intended for use by other Google Cloud client libraries that accept google.golang.org/api/option.ClientOption
// and need to inspect the configured values.
type ParsedClientOptions struct {
	Endpoint          string
	Scopes            []string
	TokenSource       oauth2.TokenSource
	Credentials       *google.Credentials // From WithCredentials
	CredentialsFile   string
	CredentialsJSON   []byte
	UserAgent         string
	APIKey            string
	Audiences         []string
	HTTPClient        *http.Client
	GRPCDialOpts      []grpc.DialOption
	GRPCConn          *grpc.ClientConn
	GRPCConnPoolSize  int
	NoAuth            bool
	TelemetryDisabled bool
	ClientCertSource  func(*tls.CertificateRequestInfo) (*tls.Certificate, error)
	Impersonation     *PublicImpersonationConfig
	UniverseDomain    string // Resolved universe domain
	Logger            *slog.Logger
	QuotaProject      string
	RequestReason     string
	AuthCredentials   *auth.Credentials // From WithAuthCredentials
}

// PublicImpersonationConfig holds publicly accessible configuration for service account impersonation.
// It mirrors the relevant parts of the internal impersonation configuration.
type PublicImpersonationConfig struct {
	TargetServiceAccount string
	Delegates            []string
}

// ParseClientOptions applies the given ClientOptions, validates them, and returns
// a ParsedClientOptions struct containing the resolved settings.
// It returns an error if the provided options are invalid according to this library's validation rules.
// This function allows external libraries to inspect configuration values
// set by users through ClientOptions without needing access to internal types.
func ParseClientOptions(opts ...option.ClientOption) (*ParsedClientOptions, error) {
	var ds internal.DialSettings
	// Apply all options to the internal DialSettings struct.
	for _, opt := range opts {
		opt.Apply(&ds)
	}

	// Validate the combined settings.
	if err := ds.Validate(); err != nil {
		return nil, err
	}

	// Populate the public ParsedClientOptions struct from the internal DialSettings.
	// Ensure copies are made for mutable types like slices to prevent external modification
	// of any shared internal state (though ds here is local, it's good practice).
	po := &ParsedClientOptions{
		Endpoint:          ds.Endpoint,
		UserAgent:         ds.UserAgent,
		APIKey:            ds.APIKey,
		CredentialsFile:   ds.CredentialsFile,
		TokenSource:       ds.TokenSource,
		Credentials:       ds.Credentials,
		HTTPClient:        ds.HTTPClient,
		GRPCConn:          ds.GRPCConn,
		GRPCConnPoolSize:  ds.GRPCConnPoolSize,
		NoAuth:            ds.NoAuth,
		TelemetryDisabled: ds.TelemetryDisabled,
		ClientCertSource:  ds.ClientCertSource,
		UniverseDomain:    ds.GetUniverseDomain(), // Uses the getter for correct precedence.
		Logger:            ds.Logger,
		QuotaProject:      ds.QuotaProject,
		RequestReason:     ds.RequestReason,
		AuthCredentials:   ds.AuthCredentials,
	}

	if len(ds.CredentialsJSON) > 0 {
		po.CredentialsJSON = make([]byte, len(ds.CredentialsJSON))
		copy(po.CredentialsJSON, ds.CredentialsJSON)
	}

	// ds.GetScopes() returns the effective scopes (user-provided or default).
	resolvedScopes := ds.GetScopes()
	if len(resolvedScopes) > 0 {
		po.Scopes = make([]string, len(resolvedScopes))
		copy(po.Scopes, resolvedScopes)
	}

	if len(ds.Audiences) > 0 { // ds.Audiences is already a copy from withAudiences.Apply
		po.Audiences = make([]string, len(ds.Audiences))
		copy(po.Audiences, ds.Audiences)
	}

	if len(ds.GRPCDialOpts) > 0 {
		po.GRPCDialOpts = make([]grpc.DialOption, len(ds.GRPCDialOpts))
		copy(po.GRPCDialOpts, ds.GRPCDialOpts)
	}

	if ds.ImpersonationConfig != nil {
		// Delegates are copied in impersonateServiceAccount.Apply
		delegatesCopy := make([]string, len(ds.ImpersonationConfig.Delegates))
		copy(delegatesCopy, ds.ImpersonationConfig.Delegates)
		po.Impersonation = &PublicImpersonationConfig{
			TargetServiceAccount: ds.ImpersonationConfig.Target,
			Delegates:            delegatesCopy,
		}
	}

	return po, nil
}
