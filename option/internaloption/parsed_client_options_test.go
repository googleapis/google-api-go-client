// Copyright 2025 Google LLC.
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
	po, err := ParseClientOptions(opts)
	if err != nil {
		t.Fatalf("ParseClientOptions(%v) err = %v, want nil", opts, err)
	}

	if po.Endpoint != testEndpoint {
		t.Errorf("po.Endpoint = %q, want %q", po.Endpoint, testEndpoint)
	}
	if po.APIKey != testAPIKey {
		t.Errorf("po.APIKey = %q, want %q", po.APIKey, testAPIKey)
	}
	if !cmp.Equal(po.Scopes, testScopes) {
		t.Errorf("po.Scopes diff (-got +want):\n%s", cmp.Diff(po.Scopes, testScopes))
	}
	if po.UserAgent != "" {
		t.Errorf("po.UserAgent != nil, got %q", po.UserAgent)
	}
}
