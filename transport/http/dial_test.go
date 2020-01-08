// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"testing"

	"google.golang.org/api/internal"
)

func TestGetEndpoint(t *testing.T) {
	testCases := []struct {
		UserEndpoint    string
		DefaultEndpoint string
		Want            string
		WantErr         bool
	}{
		{
			DefaultEndpoint: "https://foo.googleapis.com/bar/baz",
			Want:            "https://foo.googleapis.com/bar/baz",
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
			WantErr:         true,
		},
	}

	for _, tc := range testCases {
		got, err := getEndpoint(&internal.DialSettings{
			Endpoint:        tc.UserEndpoint,
			DefaultEndpoint: tc.DefaultEndpoint,
		})
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
