// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
)

func TestParseClientOptions(t *testing.T) {
	testEndpoint := "test.example.com"
	testAPIKey := "testAPIKey"
	testScopes := []string{"scope1", "scope2"}

	opts := []option.ClientOption{
		option.WithEndpoint(testEndpoint),
		option.WithAPIKey(testAPIKey),
		option.WithScopes(testScopes...),
	}

	got, err := ParseClientOptions(opts...)
	if err != nil {
		t.Fatalf("ParseClientOptions(%v) err = %v, want nil", opts, err)
	}
	if got == nil {
		t.Fatalf("ParseClientOptions(%v) got = nil, want non-nil", opts)
	}

	if got.Endpoint != testEndpoint {
		t.Errorf("ParseClientOptions().Endpoint = %q, want %q", got.Endpoint, testEndpoint)
	}
	if got.APIKey != testAPIKey {
		t.Errorf("ParseClientOptions().APIKey = %q, want %q", got.APIKey, testAPIKey)
	}
	if !cmp.Equal(got.Scopes, testScopes) {
		t.Errorf("ParseClientOptions().Scopes diff (-got +want):\n%s", cmp.Diff(got.Scopes, testScopes))
	}
	if got.UserAgent != "" { // Default UserAgent is set by internal.DialSettings if not overridden.
		// This test focuses on options passed in, not all defaults.
	}
}
