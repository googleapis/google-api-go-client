// Copyright 2026 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"net/http"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func TestNewUnsafeResolver(t *testing.T) {
	for _, tc := range []struct {
		desc                              string
		opts                              []option.ClientOption
		wantResolvedWithAPIKeyIsCustom    bool
		wantResolvedGRPCConnPoolSize      int
		wantResolvedGRPCEndpointAddress   string
		wantResolvedGRPCEndpointError     bool
		wantResolvedGRPCConnIsCustom      bool
		wantResolvedHTTPClientIsCustom    bool
		wantResolvedEnableDirectPath      bool
		wantResolvedEnableDirectPathXds   bool
		wantResolvedWithoutAuthentication bool
	}{
		{
			desc: "empty",
		},
		{
			desc: "unrelated option",
			opts: []option.ClientOption{
				option.WithUniverseDomain("foo"),
			},
		},
		{
			desc: "explicit size",
			opts: []option.ClientOption{
				option.WithGRPCConnectionPool(9),
			},
			wantResolvedGRPCConnPoolSize: 9,
		},
		{
			desc:                            "no address options",
			wantResolvedGRPCEndpointError:   false,
			wantResolvedGRPCEndpointAddress: "",
		},
		{
			desc: "basic endpoint",
			opts: []option.ClientOption{
				option.WithEndpoint("http://aaa.googleapis.com"),
			},
			wantResolvedGRPCEndpointAddress: "http://aaa.googleapis.com",
		},
		{
			desc: "custom connection",
			opts: []option.ClientOption{
				option.WithGRPCConn(new(grpc.ClientConn)),
			},
			wantResolvedGRPCConnIsCustom: true,
		},
		{
			desc: "direct path plain",
			opts: []option.ClientOption{
				EnableDirectPath(true),
			},
			wantResolvedEnableDirectPath: true,
		},
		{
			desc: "direct path xds",
			opts: []option.ClientOption{
				EnableDirectPath(true),
				EnableDirectPathXds(),
			},
			wantResolvedEnableDirectPath:    true,
			wantResolvedEnableDirectPathXds: true,
		},
		{
			desc: "api key",
			opts: []option.ClientOption{
				option.WithAPIKey("foo"),
			},
			wantResolvedWithAPIKeyIsCustom: true,
		},
		{
			desc: "without auth",
			opts: []option.ClientOption{
				option.WithoutAuthentication(),
			},
			wantResolvedWithoutAuthentication: true,
		},
		{
			desc: "http client",
			opts: []option.ClientOption{
				option.WithHTTPClient(&http.Client{}),
			},
			wantResolvedHTTPClientIsCustom: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			ur, err := NewUnsafeResolver(tc.opts...)
			if err != nil {
				t.Fatalf("NewUnsafeResolver errored: %v", err)
			}
			// check ResolvedWithAPIKeyIsCustom
			if gotCustom := ur.ResolvedWithAPIKeyIsCustom(); gotCustom != tc.wantResolvedWithAPIKeyIsCustom {
				t.Errorf("ResolvedWithAPIKeyIsCustom: got %t, want %t", gotCustom, tc.wantResolvedWithAPIKeyIsCustom)
			}
			// check ResolvedGRPCConnPoolSize
			if got := ur.ResolvedGRPCConnPoolSize(); got != tc.wantResolvedGRPCConnPoolSize {
				t.Errorf("ResolveGRPCConnPoolSize: got %d, want %d", got, tc.wantResolvedGRPCConnPoolSize)
			}
			// check ResolvedGRPCEndpoint
			gotAddr, gotErr := ur.ResolvedGRPCEndpoint()
			if tc.wantResolvedGRPCEndpointError {
				if gotErr == nil {
					t.Errorf("ResolvedGRPCEndpoint: wanted error, got success")
				}
			} else {
				if gotErr != nil {
					t.Errorf("ResolvedGRPCEndpoint: wanted success, got error %v", gotErr)
				}
			}
			if gotAddr != tc.wantResolvedGRPCEndpointAddress {
				t.Errorf("ResolvedGRPCEndpoint: address mismatch, got %q want %q", gotAddr, tc.wantResolvedGRPCEndpointAddress)
			}
			// check ResolvedGRPCConnIsCustom
			if gotCustom := ur.ResolvedGRPCConnIsCustom(); gotCustom != tc.wantResolvedGRPCConnIsCustom {
				t.Errorf("ResolvedGRPCConnIsCustom: got %t want %t", gotCustom, tc.wantResolvedGRPCConnIsCustom)
			}
			// check ResolvedHTTPClientIsCustom
			if gotCustom := ur.ResolvedHTTPClientIsCustom(); gotCustom != tc.wantResolvedHTTPClientIsCustom {
				t.Errorf("ResolvedHTTPClientIsCustom: got %t want %t", gotCustom, tc.wantResolvedHTTPClientIsCustom)
			}
			// check ResolvedEnableDirectPath
			if gotDirectPath := ur.ResolvedEnableDirectPath(); gotDirectPath != tc.wantResolvedEnableDirectPath {
				t.Errorf("ResolvedEnableDirectPath: got %t want %t", gotDirectPath, tc.wantResolvedEnableDirectPath)
			}
			if gotDirectPathXds := ur.ResolvedEnableDirectPathXds(); gotDirectPathXds != tc.wantResolvedEnableDirectPathXds {
				t.Errorf("ResolvedEnableDirectPathXds: got %t want %t", gotDirectPathXds, tc.wantResolvedEnableDirectPathXds)
			}
			// check ResolvedWithoutAuth
			if got := ur.ResolvedWithoutAuthentication(); got != tc.wantResolvedWithoutAuthentication {
				t.Errorf("ResolvedWithoutAuthentication: got %t, want %t", got, tc.wantResolvedWithoutAuthentication)
			}
		})
	}
}
