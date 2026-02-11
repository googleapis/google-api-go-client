// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/auth/grpctransport"
	"cloud.google.com/go/compute/metadata"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
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
	// Replace package var in order to assert DialContext args.
	dialContextNewAuth = func(ctx context.Context, secure bool, opts *grpctransport.Options) (grpctransport.GRPCClientConnPool, error) {
		if len(opts.GRPCDialOpts) != wantNumOpts {
			t.Fatalf("got: %d, want: %d", len(opts.GRPCDialOpts), wantNumOpts)
		}
		if opts.UniverseDomain != universeDomain {
			t.Fatalf("got: %q, want: %q", opts.UniverseDomain, universeDomain)
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
	} {
		t.Run(testcase.name, func(t *testing.T) {
			wantNumOpts = testcase.wantNumOpts
			universeDomain = testcase.ds.UniverseDomain
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
	o := &internal.DialSettings{}
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds, err := internal.Creds(context.Context(context.Background()), o)
	if err != nil {
		t.Fatalf("failed to create creds")
	}

	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	logDirectPathMisconfig(endpoint, creds.TokenSource, o)

	wantedLog := "WARNING: DirectPath is misconfigured. Please set the EnableDirectPath option along with the EnableDirectPathXds option."
	if !strings.Contains(buf.String(), wantedLog) {
		t.Fatalf("got: %v, want: %v", buf.String(), wantedLog)
	}

}

func TestLogDirectPathMisconfigWrongCredential(t *testing.T) {
	o := &internal.DialSettings{}
	o.EnableDirectPath = true
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds := &google.Credentials{}

	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	logDirectPathMisconfig(endpoint, creds.TokenSource, o)

	wantedLog := "WARNING: DirectPath is misconfigured. Please make sure the token source is fetched from GCE metadata server and the default service account is used."
	if !strings.Contains(buf.String(), wantedLog) {
		t.Fatalf("got: %v, want: %v", buf.String(), wantedLog)
	}

}

func TestLogDirectPathMisconfigNotOnGCE(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	o := &internal.DialSettings{}
	o.EnableDirectPath = true
	o.EnableDirectPathXds = true

	endpoint := "abc.googleapis.com"

	creds, err := internal.Creds(context.Context(context.Background()), o)
	if err != nil {
		t.Fatalf("failed to create creds")
	}

	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	logDirectPathMisconfig(endpoint, creds.TokenSource, o)

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

func TestIsDirectPathXdsUsed(t *testing.T) {
	tests := []struct {
		name string
		env  string
		opt  bool
		want bool
	}{
		{
			name: "Enabled by Env",
			env:  "true",
			opt:  false,
			want: true,
		},
		{
			name: "Enabled by Option",
			env:  "false",
			opt:  true,
			want: true,
		},
		{
			name: "Disabled",
			env:  "false",
			opt:  false,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("GOOGLE_CLOUD_ENABLE_DIRECT_PATH_XDS", tt.env)
			defer os.Unsetenv("GOOGLE_CLOUD_ENABLE_DIRECT_PATH_XDS")

			ds := &internal.DialSettings{EnableDirectPathXds: tt.opt}
			if got := isDirectPathXdsUsed(ds); got != tt.want {
				t.Errorf("isDirectPathXdsUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDirectPathEnabled(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		opt      bool
		endpoint string
		want     bool
	}{
		{
			name:     "Option false",
			opt:      false,
			endpoint: "dns:///foo",
			want:     false,
		},
		{
			name:     "Env disabled",
			env:      "true",
			opt:      true,
			endpoint: "dns:///foo",
			want:     false,
		},
		{
			name:     "Valid config",
			env:      "false",
			opt:      true,
			endpoint: "dns:///foo",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("GOOGLE_CLOUD_DISABLE_DIRECT_PATH", tt.env)
			defer os.Unsetenv("GOOGLE_CLOUD_DISABLE_DIRECT_PATH")

			ds := &internal.DialSettings{EnableDirectPath: tt.opt}
			if got := isDirectPathEnabled(tt.endpoint, ds); got != tt.want {
				t.Errorf("isDirectPathEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsAuthCompatible_Direct(t *testing.T) {
	ctx := context.Background()

	// Create a token that looks like it came from the GCE Metadata Server
	gceToken := (&oauth2.Token{
		AccessToken: "fake-access-token",
	}).WithExtra(map[string]interface{}{
		"oauth2.google.tokenSource":    "compute-metadata",
		"oauth2.google.serviceAccount": "default",
	})

	tests := []struct {
		name string
		ds   *internal.DialSettings
		want bool
	}{
		{
			name: "API Key - Incompatible",
			ds: &internal.DialSettings{
				APIKey: "secret-key",
			},
			want: false,
		},
		{
			name: "No Auth - Incompatible",
			ds: &internal.DialSettings{
				NoAuth: true,
			},
			want: false,
		},
		{
			name: "GCE Token - Compatible",
			ds: &internal.DialSettings{
				// Providing TokenSource here prevents internal.Creds
				// from looking for real files on your Cloudtop.
				TokenSource: &mockTokenSource{token: gceToken},
			},
			want: true,
		},
		{
			name: "Standard Token - Incompatible",
			ds: &internal.DialSettings{
				TokenSource: &mockTokenSource{
					token: &oauth2.Token{AccessToken: "regular-token"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDirectPathStatus()

			// Calling the method directly
			got := d.isAuthCompatible(ctx, tt.ds)

			if got != tt.want {
				t.Errorf("%s: isAuthCompatible() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// Helper mock for the test
type mockTokenSource struct {
	token *oauth2.Token
}

func (ts *mockTokenSource) Token() (*oauth2.Token, error) {
	return ts.token, nil
}

func TestDirectPathStatus_CheckWithReason(t *testing.T) {
	ctx := context.Background()
	detector := NewDirectPathStatus()
	// A token that satisfies the GCE Metadata requirement
	gceToken := (&oauth2.Token{AccessToken: "fake-token"}).WithExtra(map[string]interface{}{
		"oauth2.google.tokenSource":    "compute-metadata",
		"oauth2.google.serviceAccount": "default",
	})

	tests := []struct {
		name  string
		opts  []option.ClientOption
		onGCE bool
		want  string
	}{
		{
			name:  "User opted out - custom HTTP client",
			opts:  []option.ClientOption{option.WithHTTPClient(http.DefaultClient)},
			onGCE: true,
			want:  detector.UserOptedOut,
		},
		{
			name: "User opted out - DirectPathXds not enabled",
			opts: []option.ClientOption{option.WithEndpoint("dns:///google.com"),
				internaloption.EnableDirectPath(true),
			},
			onGCE: true,
			want:  detector.UserOptedOut,
		},
		{
			name: "User opted out - DirectPath not enabled",
			opts: []option.ClientOption{option.WithEndpoint("dns:///google.com"),
				internaloption.EnableDirectPathXds(),
			},
			onGCE: true,
			want:  detector.UserOptedOut,
		},
		{
			name: "User opted out - Endpoint scheme not supported",
			opts: []option.ClientOption{
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("https://google.com"), // Only dns:/// or no prefix allowed
			},
			onGCE: true,
			want:  detector.UserOptedOut,
		},
		{
			name:  "Success",
			onGCE: true, // This will override the metadata check
			opts: []option.ClientOption{
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
				option.WithTokenSource(&mockTokenSource{token: gceToken}),
			},
			want: detector.Enabled,
		},
		{
			name:  "Not on GCE",
			onGCE: false,
			opts: []option.ClientOption{
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
				option.WithTokenSource(&mockTokenSource{token: gceToken}),
			},
			want: detector.NotOnGCE,
		},
		{
			name: "Incompatible credentials - API Key",
			opts: []option.ClientOption{
				option.WithAPIKey("test-key"),
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
			},
			onGCE: true,
			want:  detector.IncompatibleCreds,
		},
		{
			name: "Incompatible credentials - NoAuth",
			opts: []option.ClientOption{
				option.WithoutAuthentication(),
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
			},
			onGCE: true,
			want:  detector.IncompatibleCreds,
		},
		{
			name: "Incompatible credentials - Token Source missing GCE metadata",
			opts: []option.ClientOption{
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
				option.WithTokenSource(&mockTokenSource{
					token: &oauth2.Token{AccessToken: "not-gce"},
				}),
			},
			onGCE: true,
			want:  detector.IncompatibleCreds,
		},
		{
			name: "Incompatible credentials - Non-default service account without permission",
			opts: []option.ClientOption{
				internaloption.EnableDirectPath(true),
				internaloption.EnableDirectPathXds(),
				option.WithEndpoint("dns:///google.com"),
				option.WithTokenSource(&mockTokenSource{
					token: (&oauth2.Token{}).WithExtra(map[string]interface{}{
						"oauth2.google.tokenSource":    "compute-metadata",
						"oauth2.google.serviceAccount": "custom-sa@project.iam.gserviceaccount.com",
					}),
				}),
			},
			onGCE: true,
			want:  detector.IncompatibleCreds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(disableDirectPath, "false")
			detector.onGCE = func() bool { return tt.onGCE }
			got := detector.CheckWithReason(ctx, tt.opts...)
			if got != tt.want {
				t.Errorf("CheckWithReason() = %v, want %v", got, tt.want)
			}
		})
	}
}
