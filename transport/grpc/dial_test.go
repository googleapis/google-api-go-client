// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"testing"
)

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
