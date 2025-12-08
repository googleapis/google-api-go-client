// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idtoken

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"google.golang.org/api/internal/credentialstype"
	"google.golang.org/api/option"
)

const (
	serviceAccountJSON = `{
		"type": "service_account",
		"project_id": "my-project"
	}`
)

func TestNewTokenSource_Validation(t *testing.T) {
	tempDir := t.TempDir()
	saFile := filepath.Join(tempDir, "sa.json")
	if err := os.WriteFile(saFile, []byte(serviceAccountJSON), 0644); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}

	ctx := context.Background()
	aud := "test-audience"

	userCreds := []byte(`{"type": "authorized_user"}`)
	externalCreds := []byte(`{"type": "external_account"}`)

	testCases := []struct {
		name           string
		opts           option.ClientOption
		wantErr        bool
		errContains    string
		errNotContains string
	}{
		{
			name:        "FileMismatch",
			opts:        option.WithAuthCredentialsFile(ExternalAccount, saFile),
			wantErr:     true,
			errContains: "credential type mismatch",
		},
		{
			name:        "JSONMismatch",
			opts:        option.WithAuthCredentialsJSON(ExternalAccount, []byte(serviceAccountJSON)),
			wantErr:     true,
			errContains: "credential type mismatch",
		},
		{
			name:           "FileCorrect",
			opts:           option.WithAuthCredentialsFile(ServiceAccount, saFile),
			wantErr:        true, // Fails later, but not with a validation error
			errNotContains: "credential type mismatch",
		},
		{
			name:        "NotAllowed",
			opts:        option.WithAuthCredentialsJSON(credentialstype.User, userCreds),
			wantErr:     true,
			errContains: "credential type not allowed",
		},
		{
			name:           "Allowed",
			opts:           option.WithAuthCredentialsJSON(credentialstype.ExternalAccount, externalCreds),
			wantErr:        true, // Fails later, but not with a validation error
			errNotContains: "credential type not allowed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewTokenSource(ctx, aud, tc.opts)

			if tc.wantErr {
				if err == nil {
					t.Fatal("got nil, want error")
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("got %q, want error containing %q", err, tc.errContains)
				}
				if tc.errNotContains != "" && strings.Contains(err.Error(), tc.errNotContains) {
					t.Errorf("got %q, want error NOT containing %q", err, tc.errNotContains)
				}
			} else if err != nil {
				t.Fatalf("got %v, want nil error", err)
			}
		})
	}
}
