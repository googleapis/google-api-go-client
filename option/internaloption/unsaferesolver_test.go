// Copyright 2026 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"testing"

	"google.golang.org/api/option"
)

func TestNewUnsafeResolver_ResolvedGRPCConnPoolSize(t *testing.T) {
	for _, tc := range []struct {
		desc string
		opts []option.ClientOption
		want int
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
			want: 9,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			ur, err := NewUnsafeResolver(tc.opts...)
			if err != nil {
				t.Fatalf("NewUnsafeResolver errored: %v", err)
			}
			got := ur.ResolvedGRPCConnPoolSize()
			if got != tc.want {
				t.Errorf("ResolveGRPCConnPoolSize: got %d, want %d", got, tc.want)
			}
		})
	}
}
