// Copyright 2017 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internal supports the options and transport packages.
package internal

import (
	"crypto/tls"
	"net/http"
	"testing"

	"cloud.google.com/go/auth"
	"google.golang.org/api/internal/impersonate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func TestSettingsValidate(t *testing.T) {
	dummyGetClientCertificate := func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) { return nil, nil }

	// Valid.
	for _, ds := range []DialSettings{
		{},
		{APIKey: "x"},
		{Scopes: []string{"s"}},
		{CredentialsFile: "f"},
		{TokenSource: dummyTS{}},
		{Credentials: &google.DefaultCredentials{}, AuthCredentials: &auth.Credentials{}}, // Support old auth to new auth automatic conversions
		{CredentialsFile: "f", AuthCredentials: &auth.Credentials{}},                      // Support old auth to new auth automatic conversions
		{CredentialsFile: "f", TokenSource: dummyTS{}},                                    // keep for backwards compatibility
		{CredentialsJSON: []byte("json")},
		{AuthCredentialsFile: "f"},
		{AuthCredentialsJSON: []byte("json")},
		{HTTPClient: &http.Client{}},
		{GRPCConn: &grpc.ClientConn{}},
		// Although NoAuth and Scopes are technically incompatible, too many
		// cloud clients add WithScopes to user-provided options to make
		// the check feasible.
		{NoAuth: true, Scopes: []string{"s"}},
		{ClientCertSource: dummyGetClientCertificate},
		{ImpersonationConfig: &impersonate.Config{Scopes: []string{"x"}}},
		{ImpersonationConfig: &impersonate.Config{}, Scopes: []string{"x"}},
	} {
		err := ds.Validate()
		if err != nil {
			t.Errorf("%+v: got %v, want nil", ds, err)
		}
	}

	// Invalid.
	for _, ds := range []DialSettings{
		{NoAuth: true, APIKey: "x"},
		{NoAuth: true, CredentialsFile: "f"},
		{NoAuth: true, AuthCredentialsFile: "f"},
		{NoAuth: true, TokenSource: dummyTS{}},
		{NoAuth: true, Credentials: &google.DefaultCredentials{}},
		{NoAuth: true, AuthCredentials: &auth.Credentials{}},
		{NoAuth: true, AuthCredentialsJSON: []byte("json")},
		{Credentials: &google.DefaultCredentials{}, CredentialsFile: "f"},
		{Credentials: &google.DefaultCredentials{}, AuthCredentialsFile: "f"},
		{Credentials: &google.DefaultCredentials{}, TokenSource: dummyTS{}},
		{Credentials: &google.DefaultCredentials{}, CredentialsJSON: []byte("json")},
		{Credentials: &google.DefaultCredentials{}, AuthCredentialsJSON: []byte("json")},
		{CredentialsFile: "f", CredentialsJSON: []byte("json")},
		{AuthCredentialsFile: "f", AuthCredentialsJSON: []byte("json")},
		{CredentialsJSON: []byte("json"), TokenSource: dummyTS{}},
		{AuthCredentialsJSON: []byte("json"), TokenSource: dummyTS{}},
		{HTTPClient: &http.Client{}, GRPCConn: &grpc.ClientConn{}},
		{HTTPClient: &http.Client{}, GRPCDialOpts: []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}},
		{Audiences: []string{"foo"}, Scopes: []string{"foo"}},
		{HTTPClient: &http.Client{}, QuotaProject: "foo"},
		{HTTPClient: &http.Client{}, RequestReason: "foo"},
		{HTTPClient: &http.Client{}, ClientCertSource: dummyGetClientCertificate},
		{ClientCertSource: dummyGetClientCertificate, GRPCConn: &grpc.ClientConn{}},
		{ClientCertSource: dummyGetClientCertificate, GRPCConnPool: struct{ ConnPool }{}},
		{ClientCertSource: dummyGetClientCertificate, GRPCDialOpts: []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}},
		{ClientCertSource: dummyGetClientCertificate, GRPCConnPoolSize: 1},
		{ImpersonationConfig: &impersonate.Config{}},
	} {
		err := ds.Validate()
		if err == nil {
			t.Errorf("%+v: got nil, want error", ds)
		}
	}

}

type dummyTS struct{}

func (dummyTS) Token() (*oauth2.Token, error) { return nil, nil }

func TestIsNewAuthLibraryEnabled(t *testing.T) {
	for _, ds := range []DialSettings{
		{EnableNewAuthLibrary: true},
		{AuthCredentials: &auth.Credentials{}},
		{AuthCredentialsJSON: []byte("json")},
		{AuthCredentialsFile: "f"},
	} {
		if !ds.IsNewAuthLibraryEnabled() {
			t.Errorf("%+v: got false, want true", ds)
		}
	}
}

func TestGetUniverseDomain(t *testing.T) {
	testCases := []struct {
		name                 string
		ds                   *DialSettings
		universeDomainEnvVar string
		want                 string
	}{
		{
			name: "none",
			ds:   &DialSettings{},
			want: "googleapis.com",
		},
		{
			name: "settings",
			ds: &DialSettings{
				UniverseDomain: "settings-example.goog",
			},
			want: "settings-example.goog",
		},
		{
			name:                 "env var",
			ds:                   &DialSettings{},
			universeDomainEnvVar: "env-example.goog",
			want:                 "env-example.goog",
		},
		{
			name: "both",
			ds: &DialSettings{
				UniverseDomain: "settings-example.goog",
			},
			universeDomainEnvVar: "env-example.goog",
			want:                 "settings-example.goog",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.universeDomainEnvVar != "" {
				t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", tc.universeDomainEnvVar)
			}

			if got := tc.ds.GetUniverseDomain(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
			if got, want := tc.ds.GetDefaultUniverseDomain(), "googleapis.com"; got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
