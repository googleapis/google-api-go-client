// Copyright 2017 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internal supports the options and transport packages.
package internal

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"google.golang.org/api/internal/cert"
	"google.golang.org/api/internal/impersonate"
	"google.golang.org/grpc"

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
		{CredentialsFile: "f", TokenSource: dummyTS{}}, // keep for backwards compatibility
		{CredentialsJSON: []byte("json")},
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
		{NoAuth: true, TokenSource: dummyTS{}},
		{NoAuth: true, Credentials: &google.DefaultCredentials{}},
		{Credentials: &google.DefaultCredentials{}, CredentialsFile: "f"},
		{Credentials: &google.DefaultCredentials{}, TokenSource: dummyTS{}},
		{Credentials: &google.DefaultCredentials{}, CredentialsJSON: []byte("json")},
		{CredentialsFile: "f", CredentialsJSON: []byte("json")},
		{CredentialsJSON: []byte("json"), TokenSource: dummyTS{}},
		{HTTPClient: &http.Client{}, GRPCConn: &grpc.ClientConn{}},
		{HTTPClient: &http.Client{}, GRPCDialOpts: []grpc.DialOption{grpc.WithInsecure()}},
		{Audiences: []string{"foo"}, Scopes: []string{"foo"}},
		{HTTPClient: &http.Client{}, QuotaProject: "foo"},
		{HTTPClient: &http.Client{}, RequestReason: "foo"},
		{HTTPClient: &http.Client{}, ClientCertSource: dummyGetClientCertificate},
		{ClientCertSource: dummyGetClientCertificate, GRPCConn: &grpc.ClientConn{}},
		{ClientCertSource: dummyGetClientCertificate, GRPCConnPool: struct{ ConnPool }{}},
		{ClientCertSource: dummyGetClientCertificate, GRPCDialOpts: []grpc.DialOption{grpc.WithInsecure()}},
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

func TestGetEndpointAndUniverse(t *testing.T) {

	fakeCertSource := func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
		return nil, fmt.Errorf("invalid source")
	}
	testCases := []struct {
		desc             string
		settings         *DialSettings
		clientCertSource cert.Source
		mtlsMode         string
		wantEnd          string
		wantUni          string
		wantErr          error
	}{
		{
			desc: "simple default",
			settings: &DialSettings{
				DefaultEndpoint: "https://foo.googleapis.com",
			},
			wantEnd: "https://foo.googleapis.com",
			wantUni: gdUniverse,
		},
		{
			desc: "simple endpoint override",
			settings: &DialSettings{
				Endpoint: "https://bar.googleapis.com",
			},
			wantEnd: "https://bar.googleapis.com",
			wantUni: gdUniverse,
		},
		{
			desc: "default + mtlsModeAuto + nocert",
			settings: &DialSettings{
				DefaultEndpoint:     "https://foo.googleapis.com",
				DefaultMTLSEndpoint: "https://foo.mtls.googleapis.com",
			},
			mtlsMode: mTLSModeAuto,
			wantEnd:  "https://foo.googleapis.com",
			wantUni:  gdUniverse,
		},
		{
			desc: "default + mtlsModeAuto + cert",
			settings: &DialSettings{
				DefaultEndpoint:     "https://foo.googleapis.com",
				DefaultMTLSEndpoint: "https://foo.mtls.googleapis.com",
			},
			clientCertSource: fakeCertSource,
			mtlsMode:         mTLSModeAuto,
			wantEnd:          "https://foo.mtls.googleapis.com",
			wantUni:          gdUniverse,
		},
		{
			desc: "default + mtlsModeAlways",
			settings: &DialSettings{
				DefaultEndpoint:     "https://foo.googleapis.com",
				DefaultMTLSEndpoint: "https://foo.mtls.googleapis.com",
			},
			mtlsMode: mTLSModeAlways,
			wantEnd:  "https://foo.mtls.googleapis.com",
			wantUni:  gdUniverse,
		},
		{
			desc: "custom uni + mtlsModeAlways",
			settings: &DialSettings{
				UniverseDomain:      "blah.com",
				DefaultEndpoint:     "https://foo.googleapis.com",
				DefaultMTLSEndpoint: "https://foo.mtls.googleapis.com",
			},
			mtlsMode: mTLSModeAlways,
			wantErr:  ErrMTLSUniverse,
		},
		{
			desc: "partial endpoint + default",
			settings: &DialSettings{
				Endpoint:        "myhost:3999",
				DefaultEndpoint: "https://foo.googleapis.com/bar/baz",
			},
			wantEnd: "https://myhost:3999/bar/baz",
			wantUni: gdUniverse,
		},
		{
			desc: "partial endpoint + default + custom uni",
			settings: &DialSettings{
				Endpoint:        "myhost:3999",
				DefaultEndpoint: "https://foo.googleapis.com/bar/baz",
				UniverseDomain:  "bar.com",
			},
			wantEnd: "https://myhost:3999/bar/baz",
			wantUni: "bar.com",
		},
		{
			desc: "partial endpoint + no default",
			settings: &DialSettings{
				Endpoint: "myhost:3999",
			},
			wantEnd: "myhost:3999",
			wantUni: gdUniverse,
		},
	}
	for _, tc := range testCases {
		mtlsMode := mTLSModeAuto
		if tc.mtlsMode != "" {
			mtlsMode = tc.mtlsMode
		}
		gotEnd, gotUni, gotErr := getEndpointAndUniverse(tc.settings, tc.clientCertSource, mtlsMode)
		if tc.wantErr != nil {
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("%q: error mismatch, got %v want %v", tc.desc, gotErr, tc.wantErr)
			}
			continue
		} else {
			if gotEnd != tc.wantEnd {
				t.Errorf("%q: endpoint mismatch, got %q want %q", tc.desc, gotEnd, tc.wantEnd)
			}
			if gotUni != tc.wantUni {
				t.Errorf("%q: universe mismatch, got %q want %q", tc.desc, gotUni, tc.wantUni)
			}
		}
	}
}
