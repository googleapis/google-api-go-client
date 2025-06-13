// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
)

type testReceiver struct {
	APIKey    string
	Endpoint  string
	Scopes    []string
	UserAgent string
}

func TestParseClientOptions(t *testing.T) {
	testEndpoint := "test.example.com"
	testAPIKey := "testAPIKey"
	testScopes := []string{"scope1", "scope2"}

	opts := []option.ClientOption{
		option.WithEndpoint(testEndpoint),
		option.WithAPIKey(testAPIKey),
		option.WithScopes(testScopes...),
	}

	receiver := &testReceiver{}
	err := ParseClientOptions(receiver, opts)
	if err != nil {
		t.Fatalf("ParseClientOptions(receiver, %v) err = %v, want nil", opts, err)
	}

	if receiver.Endpoint != testEndpoint {
		t.Errorf("receiver.Endpoint = %q, want %q", receiver.Endpoint, testEndpoint)
	}
	if receiver.APIKey != testAPIKey {
		t.Errorf("receiver.APIKey = %q, want %q", receiver.APIKey, testAPIKey)
	}
	if !cmp.Equal(receiver.Scopes, testScopes) {
		t.Errorf("receiver.Scopes diff (-got +want):\n%s", cmp.Diff(receiver.Scopes, testScopes))
	}
	if receiver.UserAgent != "" {
		t.Errorf("receiver.UserAgent != nil, got %q", receiver.UserAgent)
	}
}
