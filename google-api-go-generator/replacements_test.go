// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestDeprecatedPkgs(t *testing.T) {
	tests := []struct {
		name      string
		dp        deprecatedPkgs
		inName    string
		inVersion string
		want      string
	}{
		{
			name:      "replacement found",
			dp:        map[string]apiReplacement{"foo": {"example.com/foo", []string{"v1"}}},
			inName:    "foo",
			inVersion: "v1",
			want:      "example.com/foo",
		},
		{
			name:      "replacemet found with no versions specified",
			dp:        map[string]apiReplacement{"foo": {"example.com/foo", nil}},
			inName:    "foo",
			inVersion: "v1",
			want:      "example.com/foo",
		},
		{
			name:      "replacemet found when multiple versions present",
			dp:        map[string]apiReplacement{"foo": {"example.com/foo", []string{"v1", "v2"}}},
			inName:    "foo",
			inVersion: "v2",
			want:      "example.com/foo",
		},
		{
			name:      "no replacement found, package not in map",
			dp:        map[string]apiReplacement{"foo": {"example.com/foo", []string{"v1"}}},
			inName:    "bar",
			inVersion: "v1",
			want:      "",
		},
		{
			name:      "no replacement found, version mismatch",
			dp:        map[string]apiReplacement{"foo": {"example.com/foo", []string{"v1"}}},
			inName:    "foo",
			inVersion: "v2",
			want:      "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.dp.Get(tc.inName, tc.inVersion)
			if got != tc.want {
				t.Errorf("deprecatedPkg.Get(%v, %v): got %q, want %q", tc.inName, tc.inVersion, got, tc.want)
			}
		})
	}
}
