// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package credentialstype defines the CredType used for specifying the type of JSON credentials.
package credentialstype

import (
	"encoding/json"
	"fmt"
)

// CredType specifies the type of JSON credentials.
type CredType string

const (
	// Unknown represents an unknown JSON file type.
	Unknown CredType = ""
	// ServiceAccount represents a service account file type.
	ServiceAccount CredType = "service_account"
	// User represents a user credentials file type.
	User CredType = "authorized_user"
	// ImpersonatedServiceAccount represents an impersonated service account file type.
	//
	// IMPORTANT:
	// This credential type does not validate the credential configuration. A security
	// risk occurs when a credential configuration configured with malicious urls
	// is used.
	// You should validate credential configurations provided by untrusted sources.
	// See [Security requirements when using credential configurations from an external
	// source] https://cloud.google.com/docs/authentication/external/externally-sourced-credentials
	// for more details.
	ImpersonatedServiceAccount CredType = "impersonated_service_account"
	// ExternalAccount represents an external account file type.
	//
	// IMPORTANT:
	// This credential type does not validate the credential configuration. A security
	// risk occurs when a credential configuration configured with malicious urls
	// is used.
	// You should validate credential configurations provided by untrusted sources.
	// See [Security requirements when using credential configurations from an external
	// source] https://cloud.google.com/docs/authentication/external/externally-sourced-credentials
	// for more details.
	ExternalAccount CredType = "external_account"
)

// GetCredType returns the credentials type or the Unknown type,
// or an error for empty data or failure to unmarshal JSON.
func GetCredType(data []byte) (CredType, error) {
	var t CredType
	if len(data) == 0 {
		return t, fmt.Errorf("credential provided is 0 bytes")
	}
	var f struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &f); err != nil {
		return t, err
	}
	t = parseCredType(f.Type)
	return t, nil
}

// CheckCredentialType checks if the provided JSON bytes match the expected
// credential type. An error is returned if the JSON is invalid, the type field
// is missing, or the types do not match.
func CheckCredentialType(b []byte, want CredType) error {
	var f struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return fmt.Errorf("detect: unable to parse credential type: %w", err)
	}
	if f.Type == "" {
		return fmt.Errorf("detect: missing `type` field in credential")
	}
	got := CredType(f.Type)
	if got != want {
		return fmt.Errorf("detect: credential type mismatch: got %q, want %q", got, want)
	}
	return nil
}

func parseCredType(typeString string) CredType {
	ct := CredType(typeString)
	switch ct {
	case ServiceAccount, User, ImpersonatedServiceAccount, ExternalAccount:
		return ct
	default:
		return Unknown
	}
}
