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

	t.Run("FileMismatch", func(t *testing.T) {
		_, err := NewTokenSource(ctx, aud, option.WithAuthCredentialsFile(ExternalAccount, saFile))
		if err == nil {
			t.Fatal("got nil, want error")
		}
		if !strings.Contains(err.Error(), "credential type mismatch") {
			t.Errorf("got %q, want error containing 'credential type mismatch'", err)
		}
	})

	t.Run("JSONMismatch", func(t *testing.T) {
		_, err := NewTokenSource(ctx, aud, option.WithAuthCredentialsJSON(ExternalAccount, []byte(serviceAccountJSON)))
		if err == nil {
			t.Fatal("got nil, want error")
		}
		if !strings.Contains(err.Error(), "credential type mismatch") {
			t.Errorf("got %q, want error containing 'credential type mismatch'", err)
		}
	})

	t.Run("FileCorrect", func(t *testing.T) {
		// This should pass the validation check, but fail later since it's not a real
		// credential. We only care that it doesn't return a mismatch error.
		_, err := NewTokenSource(ctx, aud, option.WithAuthCredentialsFile(ServiceAccount, saFile))
		if err != nil && strings.Contains(err.Error(), "credential type mismatch") {
			t.Errorf("got %q, want error NOT containing 'credential type mismatch'", err)
		}
	})
}
