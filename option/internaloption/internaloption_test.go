// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/internal"
)

func TestWithCredentials(t *testing.T) {
	want := "fakeToken"
	fakeCreds := &google.Credentials{
		TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: want}),
	}
	opt := WithCredentials(fakeCreds)
	ds := &internal.DialSettings{}
	opt.Apply(ds)
	if ds.InternalCredentials == nil || ds.InternalCredentials.TokenSource == nil {
		t.Errorf("ds.InternalCredentials should be initialized")
	}
	if tok, err := ds.InternalCredentials.TokenSource.Token(); err != nil || tok.AccessToken != "fakeToken" {
		t.Errorf("ts.Token() = %q, want %q", tok.AccessToken, "")
	}
}
