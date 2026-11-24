// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/auth/grpctransport"
	"cloud.google.com/go/compute/metadata"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/time/rate"
	"google.golang.org/api/internal"
	"google.golang.org/grpc"
)

func TestDial(t *testing.T) {
	oldDialContext := dialContext
	// Replace package var in order to assert DialContext args.
	dialContext = func(ctxGot context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
		if len(opts) != 3 {
			t.Fatalf("got: %d, want: 3", len(opts))
		}
		return nil, nil
	}
	defer func() {
		dialContext = oldDialContext
	}()

	var o internal.DialSettings
	dial(context.Background(), false, &o)
}

func TestDialPoolNewAuthDialOptions(t *testing.T) {
	oldDialContextNewAuth := dialContextNewAuth
	var wantNumOpts int
	var universeDomain string
	var wantInterconnect bool
	// Replace package var in order to assert DialContext args.
	dialContextNewAuth = func(ctx context.Context, secure bool, opts *grpctransport.Options) (grpctransport.GRPCClientConnPool, error) {
		if len(opts.GRPCDialOpts) != wantNumOpts {
			t.Fatalf("got: %d, want: %d", len(opts.GRPCDialOpts), wantNumOpts)
		}
		if opts.UniverseDomain != universeDomain {
			t.Fatalf("got: %q, want: %q", opts.UniverseDomain, universeDomain)
		}
		if opts.InternalOptions == nil || opts.InternalOptions.EnableDirectPathXdsOverInterconnect != wantInterconnect {
			t.Fatalf("got EnableDirectPathXdsOverInterconnect: %v, want: %v", opts.InternalOptions.EnableDirectPathXdsOverInterconnect, wantInterconnect)
		}
		return nil, nil
	}
	defer func() {
		dialContextNewAuth = oldDialContextNewAuth
	}()

	for _, testcase := range []struct {
		name        string
		ds          *internal.DialSettings
		wantNumOpts int
	}{
		{
			name:        "no dial options",
			ds:          &internal.DialSettings{},
			wantNumOpts: 0,
		},
		{
			name: "with user agent",
			ds: &internal.DialSettings{
				UserAgent: "test",
			},
			wantNumOpts: 1,
		},
		{
			name: "universe domain",
			ds: &internal.DialSettings{
				UniverseDomain: "example.com",
			},
			wantNumOpts: 0,
		},
		{
			name: "enable direct path xds over interconnect",
			ds: &internal.DialSettings{
				EnableDirectPathXdsOverInterconnect: true,
			},
			wantNumOpts: 0,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			wantNumOpts = testcase.wantNumOpts
			universeDomain = testcase.ds.GetUniverseDomain()
			wantInterconnect = testcase.ds.EnableDirectPathXdsOverInterconnect
			dialPoolNewAuth(context.Background(), false, 1, testcase.ds)
		})
	}
}

func TestPrepareDialOptsNewAuth(t *testing.T) {
	for _, testcase := range []struct {
		name        string
		ds          *internal.DialSettings
		wantNumOpts int
	}{
		{
			name:        "empty",
			ds:          &internal.DialSettings{},
			wantNumOpts: 0,
		},
		{
			name: "user agent",
			ds: &internal.DialSettings{
				UserAgent: "test",
			},
			wantNumOpts: 1,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := prepareDialOptsNewAuth(testcase.ds)
			if len(got) != testcase.wantNumOpts {
				t.Fatalf("got %d options, want %d options", len(got), testcase.wantNumOpts)
			}
		})
	}
}

func TestCheckDirectPathEndPoint(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		endpoint string
		want     bool
	}{
		{
			name:     "empty endpoint are disallowed",
			endpoint: "",
			want:     false,
		},
		{
			name:     "dns schemes are allowed",
			endpoint: "dns:///foo",
			want:     true,
		},
		{
			name:     "host without no prefix are allowed",
			endpoint: "foo",
			want:     true,
		},
		{
			name:     "host with port are allowed",
			endpoint: "foo:1234",
			want:     true,
		},
		{
			name:     "non-dns schemes are disallowed",
			endpoint: "https://foo",
			want:     false,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			if got := checkDirectPathEndPoint(testcase.endpoint); got != testcase.want {
				t.Fatalf("got %v, want %v", got, testcase.want)
			}
		})
	}
}

func TestLogDirectPathMisconfigAttrempDirectPathNotSet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	t.Setenv(disableDirectPath, "")
	t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
	o := &internal.DialSettings{}
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds, err := internal.Creds(context.Context(context.Background()), o)
	if err != nil {
		t.Fatalf("failed to create creds")
	}

	buf := bytes.Buffer{}
	origOut := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(origOut) })

	reason, _ := canUseDirectPath(t.Context(), endpoint, creds, o)
	logDirectPathMisconfig(o, reason)

	wantedLog := "WARNING: DirectPath is misconfigured. Please set the EnableDirectPath option along with the EnableDirectPathXds option."
	if !strings.Contains(buf.String(), wantedLog) {
		t.Fatalf("got: %v, want: %v", buf.String(), wantedLog)
	}

}

func TestLogDirectPathMisconfigWrongCredential(t *testing.T) {
	t.Setenv(disableDirectPath, "")
	t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
	o := &internal.DialSettings{}
	o.EnableDirectPath = true
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds := &google.Credentials{}

	buf := bytes.Buffer{}
	origOut := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(origOut) })

	reason, _ := canUseDirectPath(t.Context(), endpoint, creds, o)
	logDirectPathMisconfig(o, reason)

	wantedLog := "WARNING: DirectPath is misconfigured. Please make sure the token source is fetched from GCE metadata server and the default service account is used."
	if !strings.Contains(buf.String(), wantedLog) {
		t.Fatalf("got: %v, want: %v", buf.String(), wantedLog)
	}

}

func TestLogDirectPathMisconfigNotOnGCE(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	t.Setenv(disableDirectPath, "")
	t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
	o := &internal.DialSettings{}
	o.EnableDirectPath = true
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds, err := internal.Creds(context.Context(context.Background()), o)
	if err != nil {
		t.Fatalf("failed to create creds")
	}

	buf := bytes.Buffer{}
	origOut := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(origOut) })

	reason, _ := canUseDirectPath(t.Context(), endpoint, creds, o)
	logDirectPathMisconfig(o, reason)

	if !metadata.OnGCE() {
		wantedLog := "WARNING: DirectPath is misconfigured. DirectPath is only available in a GCE environment."
		if !strings.Contains(buf.String(), wantedLog) {
			t.Fatalf("got: %v, want: %v", buf.String(), wantedLog)
		}
	}

}

func TestGRPCAPIKey_GetRequestMetadata(t *testing.T) {
	for _, test := range []struct {
		apiKey string
		reason string
	}{
		{
			apiKey: "MY_API_KEY",
			reason: "MY_REQUEST_REASON",
		},
	} {
		ts := grpcAPIKey{
			apiKey:        test.apiKey,
			requestReason: test.reason,
		}
		got, err := ts.GetRequestMetadata(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		want := map[string]string{
			"X-goog-api-key":        ts.apiKey,
			"X-goog-request-reason": ts.requestReason,
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}

type fakeComputeTokenSource struct {
	extra map[string]any
}

func (f fakeComputeTokenSource) Token() (*oauth2.Token, error) {
	tok := &oauth2.Token{AccessToken: "fake-token"}
	if f.extra != nil {
		tok = tok.WithExtra(f.extra)
	}
	return tok, nil
}

func defaultComputeTokenSource() oauth2.TokenSource {
	return fakeComputeTokenSource{
		extra: map[string]any{
			"oauth2.google.tokenSource":    "compute-metadata",
			"oauth2.google.serviceAccount": "default",
		},
	}
}

func nonComputeTokenSource() oauth2.TokenSource {
	return fakeComputeTokenSource{}
}

func stubGCEState(t *testing.T, gce bool, mdsUD string, mdsErr error) {
	t.Helper()
	origOnGCE := onGCE
	origMDS := metadataUniverseDomain
	onGCE = func() bool { return gce }
	metadataUniverseDomain = func(context.Context) (string, error) {
		return mdsUD, mdsErr
	}
	t.Cleanup(func() {
		onGCE = origOnGCE
		metadataUniverseDomain = origMDS
	})
}

func TestHasUniverseDomainHost(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       string
		universeDomain string
		want           bool
	}{
		{
			name:           "host_with_port",
			endpoint:       "storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "host_without_port",
			endpoint:       "storage.googleapis.com",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "exact_universe_domain_with_port",
			endpoint:       "googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "exact_universe_domain",
			endpoint:       "googleapis.com",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "sandbox_subdomain",
			endpoint:       "storage-preprod.sandbox.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "dns_scheme_with_port",
			endpoint:       "dns:///storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "dns_scheme_without_port",
			endpoint:       "dns:///storage.googleapis.com",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "uppercase_host",
			endpoint:       "STORAGE.GOOGLEAPIS.COM:443",
			universeDomain: "googleapis.com",
			want:           true,
		},
		{
			name:           "tpc_endpoint",
			endpoint:       "storage.example.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "tpc_dns_endpoint",
			endpoint:       "dns:///storage.example.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "lookalike_suffix_domain",
			endpoint:       "storage.googleapis.com.evil.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "missing_dot_boundary",
			endpoint:       "evilgoogleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "leading_dot",
			endpoint:       ".googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "empty_label",
			endpoint:       "foo..googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "userinfo_smuggling",
			endpoint:       "evil.com@storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "dns_userinfo_smuggling",
			endpoint:       "dns:///evil.com@storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "path_smuggling",
			endpoint:       "evil.com/storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "dns_custom_authority",
			endpoint:       "dns://8.8.8.8/storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "https_scheme",
			endpoint:       "https://storage.googleapis.com:443",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "query_string",
			endpoint:       "storage.googleapis.com:443?foo=bar",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "localhost",
			endpoint:       "localhost:8080",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "ipv4_loopback",
			endpoint:       "127.0.0.1:8080",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "ipv6_loopback",
			endpoint:       "[::1]:8080",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "empty_endpoint",
			endpoint:       "",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "empty_dns_target",
			endpoint:       "dns:///",
			universeDomain: "googleapis.com",
			want:           false,
		},
		{
			name:           "empty_universe_domain",
			endpoint:       "storage.googleapis.com:443",
			universeDomain: "",
			want:           false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasUniverseDomainHost(tc.endpoint, tc.universeDomain); got != tc.want {
				t.Errorf("hasUniverseDomainHost(%q, %q) = %v, want %v", tc.endpoint, tc.universeDomain, got, tc.want)
			}
		})
	}
}

func TestIsDirectPathEnabled(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		settings   *internal.DialSettings
		disableEnv string
		udEnv      string
		want       bool
	}{
		{
			name:     "gdu_default_endpoint",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     true,
		},
		{
			name:     "gdu_dns_endpoint",
			endpoint: "dns:///storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     true,
		},
		{
			name:     "gdu_explicit_universe_domain",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath: true,
				UniverseDomain:   "googleapis.com",
			},
			want: true,
		},
		{
			name:     "direct_path_disabled_in_settings",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{},
			want:     false,
		},
		{
			name:     "nil_settings",
			endpoint: "storage.googleapis.com:443",
			want:     false,
		},
		{
			name:       "disabled_by_env",
			endpoint:   "storage.googleapis.com:443",
			settings:   &internal.DialSettings{EnableDirectPath: true},
			disableEnv: "true",
			want:       false,
		},
		{
			name:     "non_gdu_universe_domain_option",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath: true,
				UniverseDomain:   "example.com",
			},
			want: false,
		},
		{
			name:     "non_gdu_universe_domain_env",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			udEnv:    "example.com",
			want:     false,
		},
		{
			name:     "non_gdu_endpoint_without_universe_domain",
			endpoint: "storage.example.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     false,
		},
		{
			name:     "non_gdu_dns_endpoint_without_universe_domain",
			endpoint: "dns:///storage.example.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     false,
		},
		{
			name:     "lookalike_endpoint",
			endpoint: "storage.googleapis.com.evil.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     false,
		},
		{
			name:     "non_dns_scheme",
			endpoint: "https://storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			want:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, tc.disableEnv)
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", tc.udEnv)

			if got := isDirectPathEnabled(tc.endpoint, tc.settings); got != tc.want {
				t.Errorf("isDirectPathEnabled(%q, %+v) = %v, want %v", tc.endpoint, tc.settings, got, tc.want)
			}
		})
	}
}

func TestCanUseDirectPath(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		settings   *internal.DialSettings
		creds      *google.Credentials
		disableEnv string
		udEnv      string
		gce        bool
		mdsUD      string
		mdsErr     error
		want       bool
		wantReason string
	}{
		{
			name:     "gdu_gce_adc",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			gce:  true,
			want: true,
		},
		{
			name:     "gdu_gce_adc_with_trailing_newline",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com\n", nil
				},
			},
			gce:  true,
			want: true,
		},
		{
			name:       "direct_path_disabled",
			endpoint:   "storage.googleapis.com:443",
			settings:   &internal.DialSettings{},
			creds:      &google.Credentials{TokenSource: defaultComputeTokenSource()},
			gce:        true,
			want:       false,
			wantReason: directPathNotEnabledReason,
		},
		{
			name:       "disabled_by_env",
			endpoint:   "storage.googleapis.com:443",
			settings:   &internal.DialSettings{EnableDirectPath: true},
			creds:      &google.Credentials{TokenSource: defaultComputeTokenSource()},
			disableEnv: "true",
			gce:        true,
			want:       false,
		},
		{
			name:     "non_dns_scheme_no_warning",
			endpoint: "https://storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds:    &google.Credentials{TokenSource: defaultComputeTokenSource()},
			gce:      true,
			want:     false,
		},
		{
			name:     "tpc_gce_adc_with_trailing_newline",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "example.com\n", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "tpc_gce_adc_without_client_universe_domain",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "example.com", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "creds_universe_domain_error_fails_closed",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "", errors.New("mds unavailable")
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "token_source_without_provider_blocked_by_tpc_mds",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
			},
			gce:        true,
			mdsUD:      "example.com",
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "token_source_without_provider_mds_error_fails_closed",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
			},
			gce:        true,
			mdsErr:     errors.New("mds error"),
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "token_source_without_provider_allowed_by_gdu_mds",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
			},
			gce:   true,
			mdsUD: "googleapis.com",
			want:  true,
		},
		{
			name:     "tpc_endpoint_without_client_universe_domain",
			endpoint: "storage.example.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "tpc_client_universe_domain_with_gdu_endpoint",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath: true,
				UniverseDomain:   "example.com",
			},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "tpc_env_var_with_gdu_endpoint",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			udEnv:      "example.com",
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "not_on_gce",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			gce:        false,
			want:       false,
			wantReason: directPathNotOnGCEReason,
		},
		{
			name:       "nil_creds",
			endpoint:   "storage.googleapis.com:443",
			settings:   &internal.DialSettings{EnableDirectPath: true},
			gce:        true,
			want:       false,
			wantReason: directPathIncompatibleCredReason,
		},
		{
			name:       "creds_with_nil_token_source",
			endpoint:   "storage.googleapis.com:443",
			settings:   &internal.DialSettings{EnableDirectPath: true},
			creds:      &google.Credentials{},
			gce:        true,
			want:       false,
			wantReason: directPathIncompatibleCredReason,
		},
		{
			name:     "non_compute_token_source_disallowed",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{EnableDirectPath: true},
			creds: &google.Credentials{
				TokenSource: nonComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathIncompatibleCredReason,
		},
		{
			name:     "allow_non_default_sa_tpc_creds_rejected",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:              true,
				AllowNonDefaultServiceAccount: true,
			},
			creds: &google.Credentials{
				TokenSource: nonComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "example.com", nil
				},
			},
			gce:        true,
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "allow_non_default_sa_gdu_creds_on_tpc_vm_rejected",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:              true,
				AllowNonDefaultServiceAccount: true,
			},
			creds: &google.Credentials{
				TokenSource: nonComputeTokenSource(),
			},
			gce:        true,
			mdsUD:      "example.com",
			want:       false,
			wantReason: directPathUnsupportedUniverseReason,
		},
		{
			name:     "allow_non_default_sa_gdu_creds_on_gdu_vm_allowed",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:              true,
				AllowNonDefaultServiceAccount: true,
			},
			creds: &google.Credentials{
				TokenSource: nonComputeTokenSource(),
			},
			gce:   true,
			mdsUD: "googleapis.com",
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, tc.disableEnv)
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", tc.udEnv)
			stubGCEState(t, tc.gce, tc.mdsUD, tc.mdsErr)

			gotReason, got := canUseDirectPath(t.Context(), tc.endpoint, tc.creds, tc.settings)
			if got != tc.want || gotReason != tc.wantReason {
				t.Errorf("canUseDirectPath(%q) = (%q, %v), want (%q, %v)", tc.endpoint, gotReason, got, tc.wantReason, tc.want)
			}
		})
	}
}

func TestDial_DirectPathUniverseDomainGating(t *testing.T) {
	tests := []struct {
		name        string
		settings    *internal.DialSettings
		udEnv       string
		mdsUD       string
		mdsErr      error
		wantTarget  string
		wantNumOpts int
	}{
		{
			name: "gdu_gce_xds_rewrites_to_google_c2p",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "googleapis.com", nil
					},
				},
			},
			wantTarget:  "google-c2p:///storage.googleapis.com",
			wantNumOpts: 1,
		},
		{
			name: "gdu_gce_xds_dns_scheme_with_port_rewrites_to_google_c2p",
			settings: &internal.DialSettings{
				Endpoint:            "dns:///storage.googleapis.com:443",
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
				TelemetryDisabled:   true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "googleapis.com", nil
					},
				},
			},
			wantTarget:  "google-c2p:///storage.googleapis.com",
			wantNumOpts: 1,
		},
		{
			name: "gdu_gce_xds_dns_scheme_without_port_rewrites_to_google_c2p",
			settings: &internal.DialSettings{
				Endpoint:            "dns:///storage.googleapis.com",
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
				TelemetryDisabled:   true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "googleapis.com", nil
					},
				},
			},
			wantTarget:  "google-c2p:///storage.googleapis.com",
			wantNumOpts: 1,
		},
		{
			name: "gdu_gce_grpclb_rewrites_to_dns",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "googleapis.com", nil
					},
				},
			},
			wantTarget:  "dns:///storage.googleapis.com:443",
			wantNumOpts: 3,
		},
		{
			name: "tpc_universe_domain_option_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				UniverseDomain:          "example.com",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "example.com", nil
					},
				},
			},
			wantTarget:  "storage.example.com:443",
			wantNumOpts: 2,
		},
		{
			name: "tpc_env_var_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "example.com", nil
					},
				},
			},
			udEnv:       "example.com",
			wantTarget:  "storage.example.com:443",
			wantNumOpts: 2,
		},
		{
			name: "tpc_endpoint_without_universe_domain_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				Endpoint:                "storage.example.com:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "googleapis.com", nil
					},
				},
			},
			wantTarget:  "storage.example.com:443",
			wantNumOpts: 2,
		},
		{
			name: "tpc_creds_without_client_universe_domain_blocks_xds_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "example.com", nil
					},
				},
			},
			wantTarget:  "storage.googleapis.com:443",
			wantNumOpts: 2,
		},
		{
			name: "tpc_creds_without_client_universe_domain_blocks_grpclb_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "example.com", nil
					},
				},
			},
			wantTarget:  "storage.googleapis.com:443",
			wantNumOpts: 2,
		},
		{
			name: "creds_universe_domain_error_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				Credentials: &google.Credentials{
					TokenSource: defaultComputeTokenSource(),
					UniverseDomainProvider: func() (string, error) {
						return "", errors.New("mds timeout")
					},
				},
			},
			wantTarget:  "storage.googleapis.com:443",
			wantNumOpts: 2,
		},
		{
			name: "token_source_on_tpc_vm_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				TokenSource:             defaultComputeTokenSource(),
			},
			mdsUD:       "example.com",
			wantTarget:  "storage.googleapis.com:443",
			wantNumOpts: 2,
		},
		{
			name: "token_source_with_mds_error_blocks_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				TokenSource:             defaultComputeTokenSource(),
			},
			mdsErr:      errors.New("mds error"),
			wantTarget:  "storage.googleapis.com:443",
			wantNumOpts: 2,
		},
		{
			name: "token_source_on_gdu_vm_enables_directpath",
			settings: &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				TokenSource:             defaultComputeTokenSource(),
			},
			mdsUD:       "googleapis.com",
			wantTarget:  "google-c2p:///storage.googleapis.com",
			wantNumOpts: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, "")
			t.Setenv(enableDirectPathXds, "")
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", tc.udEnv)
			t.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "false")
			t.Setenv("EXPERIMENTAL_GOOGLE_API_USE_S2A", "false")
			stubGCEState(t, true, tc.mdsUD, tc.mdsErr)

			origDialContext := dialContext
			origTimeoutDialer := timeoutDialerOption
			timeoutDialerOption = nil
			var gotTarget string
			var gotNumOpts int
			dialContext = func(_ context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
				gotTarget = target
				gotNumOpts = len(opts)
				return nil, nil
			}
			t.Cleanup(func() {
				dialContext = origDialContext
				timeoutDialerOption = origTimeoutDialer
			})

			if _, err := dial(t.Context(), false, tc.settings); err != nil {
				t.Fatalf("dial() unexpected error: %v", err)
			}
			if gotTarget != tc.wantTarget {
				t.Errorf("dial() target = %q, want %q", gotTarget, tc.wantTarget)
			}
			if gotNumOpts != tc.wantNumOpts {
				t.Errorf("dial() len(opts) = %d, want %d", gotNumOpts, tc.wantNumOpts)
			}
		})
	}
}

func TestFetchMetadataUniverseDomain(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		want       string
		wantErr    bool
	}{
		{
			name:       "gdu_response",
			statusCode: http.StatusOK,
			body:       "googleapis.com",
			want:       "googleapis.com",
		},
		{
			name:       "tpc_response_trimmed",
			statusCode: http.StatusOK,
			body:       "example.com\n",
			want:       "example.com",
		},
		{
			name:       "empty_response_defaults_to_gdu",
			statusCode: http.StatusOK,
			body:       "",
			want:       "googleapis.com",
		},
		{
			name:       "not_found_defaults_to_gdu",
			statusCode: http.StatusNotFound,
			body:       "not found",
			want:       "googleapis.com",
		},
		{
			name:       "non_404_error_fails_closed",
			statusCode: http.StatusForbidden,
			body:       "forbidden",
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/computeMetadata/v1/universe/universe-domain" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.WriteHeader(tc.statusCode)
				_, _ = w.Write([]byte(tc.body))
			}))
			t.Cleanup(srv.Close)

			t.Setenv("GCE_METADATA_HOST", strings.TrimPrefix(srv.URL, "http://"))

			got, err := fetchMetadataUniverseDomain(t.Context())
			if (err != nil) != tc.wantErr {
				t.Fatalf("fetchMetadataUniverseDomain() error = %v, wantErr %v", err, tc.wantErr)
			}
			if err != nil {
				return
			}
			if got != tc.want {
				t.Errorf("fetchMetadataUniverseDomain() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestLogDirectPathMisconfig_NonDefaultUniverseDomain(t *testing.T) {
	const wantWarning = "WARNING: DirectPath is disabled. Non-default universe domain is not supported."

	tests := []struct {
		name     string
		endpoint string
		settings *internal.DialSettings
		creds    *google.Credentials
		mdsUD    string
		mdsErr   error
	}{
		{
			name:     "non_gdu_universe_domain_option",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
				UniverseDomain:      "example.com",
			},
			creds: &google.Credentials{TokenSource: defaultComputeTokenSource()},
		},
		{
			name:     "non_gdu_endpoint_host",
			endpoint: "storage.example.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			},
			creds: &google.Credentials{TokenSource: defaultComputeTokenSource()},
		},
		{
			name:     "non_gdu_credentials_universe_domain",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "example.com", nil
				},
			},
		},
		{
			name:     "credentials_universe_domain_error",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			},
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "", errors.New("mds error")
				},
			},
		},
		{
			name:     "non_gdu_mds_without_provider",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			},
			creds: &google.Credentials{TokenSource: defaultComputeTokenSource()},
			mdsUD: "example.com",
		},
		{
			name:     "mds_error_without_provider",
			endpoint: "storage.googleapis.com:443",
			settings: &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			},
			creds:  &google.Credentials{TokenSource: defaultComputeTokenSource()},
			mdsErr: errors.New("mds error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, "")
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
			stubGCEState(t, true, tc.mdsUD, tc.mdsErr)

			var buf bytes.Buffer
			origOut := log.Writer()
			log.SetOutput(&buf)
			t.Cleanup(func() { log.SetOutput(origOut) })

			reason, _ := canUseDirectPath(t.Context(), tc.endpoint, tc.creds, tc.settings)
			logDirectPathMisconfig(tc.settings, reason)

			if !strings.Contains(buf.String(), wantWarning) {
				t.Errorf("logDirectPathMisconfig(%q) logged %q, want substring %q", tc.endpoint, buf.String(), wantWarning)
			}
		})
	}
}

func TestLogDirectPathMisconfig(t *testing.T) {
	tests := []struct {
		name     string
		settings *internal.DialSettings
		xdsEnv   string
		reason   string
		want     string
	}{
		{
			name:     "xds_option_logs_warning",
			settings: &internal.DialSettings{EnableDirectPath: true, EnableDirectPathXds: true},
			reason:   directPathUnsupportedUniverseReason,
			want:     "WARNING: " + directPathUnsupportedUniverseReason,
		},
		{
			name:     "xds_env_logs_warning",
			settings: &internal.DialSettings{EnableDirectPath: true},
			xdsEnv:   "true",
			reason:   directPathNotOnGCEReason,
			want:     "WARNING: " + directPathNotOnGCEReason,
		},
		{
			name:     "xds_disabled_logs_nothing",
			settings: &internal.DialSettings{EnableDirectPath: true},
			reason:   directPathUnsupportedUniverseReason,
		},
		{
			name:     "empty_reason_logs_nothing",
			settings: &internal.DialSettings{EnableDirectPath: true, EnableDirectPathXds: true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(enableDirectPathXds, tc.xdsEnv)

			var buf bytes.Buffer
			origOut := log.Writer()
			log.SetOutput(&buf)
			t.Cleanup(func() { log.SetOutput(origOut) })

			logDirectPathMisconfig(tc.settings, tc.reason)

			if tc.want == "" {
				if buf.Len() != 0 {
					t.Errorf("logDirectPathMisconfig(%q) logged %q, want empty output", tc.reason, buf.String())
				}
				return
			}
			if !strings.Contains(buf.String(), tc.want) {
				t.Errorf("logDirectPathMisconfig(%q) logged %q, want substring %q", tc.reason, buf.String(), tc.want)
			}
		})
	}
}

func TestLogDirectPathMisconfig_ValidGDU(t *testing.T) {
	tests := []struct {
		name  string
		creds *google.Credentials
		mdsUD string
	}{
		{
			name: "with_universe_domain_provider",
			creds: &google.Credentials{
				TokenSource: defaultComputeTokenSource(),
				UniverseDomainProvider: func() (string, error) {
					return "googleapis.com\n", nil
				},
			},
		},
		{
			name:  "without_provider_gdu_mds",
			creds: &google.Credentials{TokenSource: defaultComputeTokenSource()},
			mdsUD: "googleapis.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, "")
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
			stubGCEState(t, true, tc.mdsUD, nil)

			var buf bytes.Buffer
			origOut := log.Writer()
			log.SetOutput(&buf)
			t.Cleanup(func() { log.SetOutput(origOut) })

			settings := &internal.DialSettings{
				EnableDirectPath:    true,
				EnableDirectPathXds: true,
			}
			reason, _ := canUseDirectPath(t.Context(), "storage.googleapis.com:443", tc.creds, settings)
			logDirectPathMisconfig(settings, reason)

			if buf.Len() != 0 {
				t.Errorf("logDirectPathMisconfig() logged %q, want empty output", buf.String())
			}
		})
	}
}

type countingTokenSource struct {
	base  oauth2.TokenSource
	calls int
}

func (c *countingTokenSource) Token() (*oauth2.Token, error) {
	c.calls++
	return c.base.Token()
}

func TestDial_DirectPathDoesNotDuplicateLookups(t *testing.T) {
	tests := []struct {
		name        string
		mdsUD       string
		wantWarning string
	}{
		{
			name:  "healthy_gdu",
			mdsUD: "googleapis.com",
		},
		{
			name:        "rejected_non_gdu_mds",
			mdsUD:       "example.com",
			wantWarning: "WARNING: " + directPathUnsupportedUniverseReason,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(disableDirectPath, "")
			t.Setenv(enableDirectPathXds, "")
			t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
			t.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "false")
			t.Setenv("EXPERIMENTAL_GOOGLE_API_USE_S2A", "false")

			origOnGCE := onGCE
			origMDS := metadataUniverseDomain
			origDialContext := dialContext
			origOut := log.Writer()
			onGCE = func() bool { return true }
			mdsCalls := 0
			metadataUniverseDomain = func(context.Context) (string, error) {
				mdsCalls++
				return tc.mdsUD, nil
			}
			dialContext = func(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error) {
				return nil, nil
			}
			var buf bytes.Buffer
			log.SetOutput(&buf)
			logRateLimiter = rate.Sometimes{Interval: time.Second}
			t.Cleanup(func() {
				onGCE = origOnGCE
				metadataUniverseDomain = origMDS
				dialContext = origDialContext
				log.SetOutput(origOut)
				logRateLimiter = rate.Sometimes{Interval: time.Second}
			})

			ts := &countingTokenSource{base: defaultComputeTokenSource()}
			settings := &internal.DialSettings{
				DefaultEndpointTemplate: "storage.UNIVERSE_DOMAIN:443",
				EnableDirectPath:        true,
				EnableDirectPathXds:     true,
				TelemetryDisabled:       true,
				TokenSource:             ts,
			}

			if _, err := dial(t.Context(), false, settings); err != nil {
				t.Fatalf("dial() unexpected error: %v", err)
			}
			if ts.calls != 1 {
				t.Errorf("dial() token lookups = %d, want 1", ts.calls)
			}
			if mdsCalls != 1 {
				t.Errorf("dial() metadata universe domain lookups = %d, want 1", mdsCalls)
			}
			if tc.wantWarning == "" {
				if buf.Len() != 0 {
					t.Errorf("dial() logged %q, want empty output", buf.String())
				}
			} else if !strings.Contains(buf.String(), tc.wantWarning) {
				t.Errorf("dial() logged %q, want substring %q", buf.String(), tc.wantWarning)
			}
		})
	}
}

func TestFetchMetadataUniverseDomain_Timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	t.Cleanup(srv.Close)

	origTimeout := metadataUniverseDomainTimeout
	metadataUniverseDomainTimeout = 50 * time.Millisecond
	t.Cleanup(func() { metadataUniverseDomainTimeout = origTimeout })

	t.Setenv("GCE_METADATA_HOST", strings.TrimPrefix(srv.URL, "http://"))

	got, err := fetchMetadataUniverseDomain(t.Context())
	if err == nil {
		t.Fatalf("fetchMetadataUniverseDomain() = (%q, nil), want timeout error", got)
	}
}
