// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"google.golang.org/api/internal/cert"
)

const (
	testMTLSEndpoint     = "test.mtls.endpoint"
	testRegularEndpoint  = "test.endpoint"
	testOverrideEndpoint = "test.override.endpoint"
)

var dummyClientCertSource = func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) { return nil, nil }

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

func TestGetEndpointWithClientCertSource(t *testing.T) {

	testCases := []struct {
		UserEndpoint        string
		DefaultEndpoint     string
		DefaultMTLSEndpoint string
		Want                string
		WantErr             bool
	}{
		{
			DefaultEndpoint:     "https://foo.googleapis.com/bar/baz",
			DefaultMTLSEndpoint: "https://foo.mtls.googleapis.com/bar/baz",
			Want:                "https://foo.mtls.googleapis.com/bar/baz",
		},
		{
			DefaultEndpoint:     "https://staging-foo.sandbox.googleapis.com/bar/baz",
			DefaultMTLSEndpoint: "https://staging-foo.mtls.sandbox.googleapis.com/bar/baz",
			Want:                "https://staging-foo.mtls.sandbox.googleapis.com/bar/baz",
		},
		{
			UserEndpoint:    "myhost:3999",
			DefaultEndpoint: "https://foo.googleapis.com/bar/baz",
			Want:            "https://myhost:3999/bar/baz",
		},
		{
			UserEndpoint:    "https://host/path/to/bar",
			DefaultEndpoint: "https://foo.googleapis.com/bar/baz",
			Want:            "https://host/path/to/bar",
		},
		{
			UserEndpoint:    "host:port",
			DefaultEndpoint: "",
			Want:            "host:port",
		},
	}

	for _, tc := range testCases {
		got, _, err := getEndpointAndUniverse(&DialSettings{
			Endpoint:            tc.UserEndpoint,
			DefaultEndpoint:     tc.DefaultEndpoint,
			DefaultMTLSEndpoint: tc.DefaultMTLSEndpoint,
		}, dummyClientCertSource, getMTLSMode())
		if tc.WantErr && err == nil {
			t.Errorf("want err, got nil err")
			continue
		}
		if !tc.WantErr && err != nil {
			t.Errorf("want nil err, got %v", err)
			continue
		}
		if tc.Want != got {
			t.Errorf("getEndpoint(%q, %q): got %v; want %v", tc.UserEndpoint, tc.DefaultEndpoint, got, tc.Want)
		}
	}
}

func TestGetGRPCTransportConfigAndEndpoint(t *testing.T) {
	testCases := []struct {
		Desc          string
		InputSettings *DialSettings
		S2ARespFunc   func() (string, error)
		MTLSEnabled   func() bool
		WantEndpoint  string
	}{
		{
			"no client cert, endpoint is MTLS enabled, S2A address not empty",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testMTLSEndpoint,
		},
		{
			"has client cert",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
				ClientCertSource:    dummyClientCertSource,
			},
			validConfigResp,
			func() bool { return true },
			testMTLSEndpoint,
		},
		{
			"no client cert, endpoint is not MTLS enabled",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			validConfigResp,
			func() bool { return false },
			testRegularEndpoint,
		},
		{
			"no client cert, endpoint is MTLS enabled, S2A address empty",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			invalidConfigResp,
			func() bool { return true },
			testRegularEndpoint,
		},
		{
			"no client cert, endpoint is MTLS enabled, S2A address not empty, override endpoint",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
				Endpoint:            testOverrideEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testOverrideEndpoint,
		},
	}
	defer setupTest()()

	for _, tc := range testCases {
		httpGetMetadataMTLSConfig = tc.S2ARespFunc
		mtlsEndpointEnabledForS2A = tc.MTLSEnabled
		if tc.InputSettings.ClientCertSource != nil {
			os.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "true")
		} else {
			os.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "false")
		}
		_, endpoint, _ := GetGRPCTransportConfigAndEndpoint(tc.InputSettings)
		if tc.WantEndpoint != endpoint {
			t.Errorf("%s: want endpoint: [%s], got [%s]", tc.Desc, tc.WantEndpoint, endpoint)
		}
		// Let the cached MTLS config expire at the end of each test case.
		time.Sleep(2 * time.Millisecond)
	}
}

func TestGetHTTPTransportConfigAndEndpoint(t *testing.T) {
	testCases := []struct {
		Desc          string
		InputSettings *DialSettings
		S2ARespFunc   func() (string, error)
		MTLSEnabled   func() bool
		WantEndpoint  string
		DialFuncNil   bool
	}{
		{
			"no client cert, endpoint is MTLS enabled, S2A address not empty",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testMTLSEndpoint,
			false,
		},
		{
			"has client cert",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
				ClientCertSource:    dummyClientCertSource,
			},
			validConfigResp,
			func() bool { return true },
			testMTLSEndpoint,
			true,
		},
		{
			"no client cert, endpoint is not MTLS enabled",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			validConfigResp,
			func() bool { return false },
			testRegularEndpoint,
			true,
		},
		{
			"no client cert, endpoint is MTLS enabled, S2A address empty",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
			},
			invalidConfigResp,
			func() bool { return true },
			testRegularEndpoint,
			true,
		},
		{
			"no client cert, endpoint is MTLS enabled, S2A address not empty, override endpoint",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
				Endpoint:            testOverrideEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testOverrideEndpoint,
			false,
		},
		{
			"no client cert, S2A address not empty, but DefaultMTLSEndpoint is not set",
			&DialSettings{
				DefaultMTLSEndpoint: "",
				DefaultEndpoint:     testRegularEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testRegularEndpoint,
			true,
		},
		{
			"no client cert, S2A address not empty, override endpoint is set",
			&DialSettings{
				DefaultMTLSEndpoint: "",
				Endpoint:            testOverrideEndpoint,
			},
			validConfigResp,
			func() bool { return true },
			testOverrideEndpoint,
			false,
		},
		{
			"no client cert, endpoint is MTLS enabled, S2A address not empty, custom HTTP client",
			&DialSettings{
				DefaultMTLSEndpoint: testMTLSEndpoint,
				DefaultEndpoint:     testRegularEndpoint,
				HTTPClient:          http.DefaultClient,
			},
			validConfigResp,
			func() bool { return true },
			testRegularEndpoint,
			true,
		},
	}

	defer setupTest()()

	for _, tc := range testCases {
		httpGetMetadataMTLSConfig = tc.S2ARespFunc
		mtlsEndpointEnabledForS2A = tc.MTLSEnabled
		if tc.InputSettings.ClientCertSource != nil {
			os.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "true")
		} else {
			os.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "false")
		}
		_, dialFunc, endpoint, _ := GetHTTPTransportConfigAndEndpoint(tc.InputSettings)
		if tc.WantEndpoint != endpoint {
			t.Errorf("%s: want endpoint: [%s], got [%s]", tc.Desc, tc.WantEndpoint, endpoint)
		}
		if want, got := tc.DialFuncNil, dialFunc == nil; want != got {
			t.Errorf("%s: expecting returned dialFunc is nil: [%v], got [%v]", tc.Desc, tc.DialFuncNil, got)
		}
		// Let MTLS config expire at end of each test case.
		time.Sleep(2 * time.Millisecond)
	}
}

func setupTest() func() {
	oldDefaultMTLSEnabled := mtlsEndpointEnabledForS2A
	oldHTTPGet := httpGetMetadataMTLSConfig
	oldExpiry := configExpiry
	oldUseS2A := os.Getenv(googleAPIUseS2AEnv)
	oldUseClientCert := os.Getenv("GOOGLE_API_USE_CLIENT_CERTIFICATE")

	configExpiry = time.Millisecond
	os.Setenv(googleAPIUseS2AEnv, "true")

	return func() {
		httpGetMetadataMTLSConfig = oldHTTPGet
		mtlsEndpointEnabledForS2A = oldDefaultMTLSEnabled
		configExpiry = oldExpiry
		os.Setenv(googleAPIUseS2AEnv, oldUseS2A)
		os.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", oldUseClientCert)
	}
}
