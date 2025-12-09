// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package credentialstype

import (
	"strings"
	"testing"
)

func TestCheckCredentialType(t *testing.T) {
	serviceAccountJSON := []byte(`{"type": "service_account"}`)
	invalidJSON := []byte(`invalid json`)
	missingTypeJSON := []byte(`{"foo": "bar"}`)
	emptyJSON := []byte(``)

	testCases := []struct {
		name        string
		jsonBytes   []byte
		expected    CredType
		allowed     []CredType
		wantErr     bool
		errContains string
	}{
		{
			name:        "ExpectedMismatch",
			jsonBytes:   serviceAccountJSON,
			expected:    AuthorizedUser,
			wantErr:     true,
			errContains: "credential type mismatch: got \"service_account\", expected \"authorized_user\"",
		},
		{
			name:        "AllowedMismatch",
			jsonBytes:   serviceAccountJSON,
			expected:    ServiceAccount,
			allowed:     []CredType{AuthorizedUser, ExternalAccount},
			wantErr:     true,
			errContains: "credential type not allowed: \"service_account\"",
		},
		{
			name:      "AllowedSuccess",
			jsonBytes: serviceAccountJSON,
			expected:  ServiceAccount,
			allowed:   []CredType{ServiceAccount, AuthorizedUser},
			wantErr:   false,
		}, {
			name:      "NoAllowedSpecified",
			jsonBytes: serviceAccountJSON,
			expected:  ServiceAccount,
			allowed:   nil,
			wantErr:   false,
		},
		{
			name:        "InvalidJSON",
			jsonBytes:   invalidJSON,
			expected:    ServiceAccount,
			wantErr:     true,
			errContains: "unable to parse credential type",
		},
		{
			name:        "MissingTypeField",
			jsonBytes:   missingTypeJSON,
			expected:    ServiceAccount,
			wantErr:     true,
			errContains: "missing `type` field in credential",
		},
		{
			name:        "EmptyJSON",
			jsonBytes:   emptyJSON,
			expected:    ServiceAccount,
			wantErr:     true,
			errContains: "unable to parse credential type",
		},
		{
			name:      "GetCredType_Success",
			jsonBytes: serviceAccountJSON,
			expected:  ServiceAccount,
			allowed:   nil,
			wantErr:   false,
		},
		{
			name:        "GetCredType_EmptyJSON",
			jsonBytes:   emptyJSON,
			expected:    Unknown,
			wantErr:     true,
			errContains: "credential provided is 0 bytes",
		},
		{
			name:        "GetCredType_InvalidJSON",
			jsonBytes:   invalidJSON,
			expected:    Unknown,
			wantErr:     true,
			errContains: "invalid character",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Testing GetCredType separately as it's also a public function
			if strings.HasPrefix(tc.name, "GetCredType_") {
				gotType, err := GetCredType(tc.jsonBytes)
				if tc.wantErr {
					if err == nil {
						t.Fatalf("GetCredType() = nil, want error for %q", tc.name)
					}
					if !strings.Contains(err.Error(), tc.errContains) {
						t.Errorf("GetCredType() error = %q, want error containing %q for %q", err, tc.errContains, tc.name)
					}
				} else {
					if err != nil {
						t.Fatalf("GetCredType() = %v, want nil error for %q", err, tc.name)
					}
					if gotType != tc.expected {
						t.Errorf("GetCredType() got type = %q, want %q for %q", gotType, tc.expected, tc.name)
					}
				}
				return
			}

			// Test CheckCredentialType
			err := CheckCredentialType(tc.jsonBytes, tc.expected, tc.allowed...)
			if tc.wantErr {
				if err == nil {
					t.Fatal("got nil, want error")
				}
				if !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("got error = %q, want error containing %q", err, tc.errContains)
				}
			} else if err != nil {
				t.Fatalf("got error = %v, want nil error", err)
			}
		})
	}
}

func TestParseCredType(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected CredType
	}{
		{
			name:     "ServiceAccount",
			input:    "service_account",
			expected: ServiceAccount,
		},
		{
			name:     "AuthorizedUser",
			input:    "authorized_user",
			expected: AuthorizedUser,
		},
		{
			name:     "ImpersonatedServiceAccount",
			input:    "impersonated_service_account",
			expected: ImpersonatedServiceAccount,
		},
		{
			name:     "ExternalAccount",
			input:    "external_account",
			expected: ExternalAccount,
		},
		{
			name:     "GDCHServiceAccount",
			input:    "gdc_service_account",
			expected: GDCHServiceAccount,
		},
		{
			name:     "ExternalAccountAuthorizedUser",
			input:    "external_account_authorized_user",
			expected: ExternalAccountAuthorizedUser,
		},
		{
			name:     "UnknownType",
			input:    "not_a_real_type",
			expected: Unknown,
		},
		{
			name:     "EmptyString",
			input:    "",
			expected: Unknown,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseCredType(tc.input)
			if got != tc.expected {
				t.Errorf("parseCredType(%q) got %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}
