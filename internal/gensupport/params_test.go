// Copyright 2021 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"testing"

	"google.golang.org/api/googleapi"
)

func TestSetOptionsGetMulti(t *testing.T) {
	co := googleapi.QueryParameter("key", "foo", "bar")
	urlParams := make(URLParams)
	SetOptions(urlParams, co)
	if got, want := urlParams.Encode(), "key=foo&key=bar"; got != want {
		t.Fatalf("URLParams.Encode() = %q, want %q", got, want)
	}
}
