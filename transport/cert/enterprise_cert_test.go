// Copyright 2022 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cert

import (
	"testing"
)

func TestEnterpriseCertificateProxySource_ConfigMissing(t *testing.T) {
	source, err := NewEnterpriseCertificateProxySource("missing.json")
	if err != nil {
		t.Fatal("NewEnterpriseCertificateProxySource: expected nil error returned when config is missing.")
	}
	if source != nil {
		t.Error("NewEnterpriseCertificateProxySource: expected nil source returned when config is missing.")
	}
}

// This test launches a mock signer binary "test_signer.go" that uses a valid pem file.
func TestEnterpriseCertificateProxySource_GetClientCertificateSuccess(t *testing.T) {
	source, err := NewEnterpriseCertificateProxySource("testdata/enterprise_certificate_config.json")
	if err != nil {
		t.Fatal(err)
	}
	cert, err := source(nil)
	if err != nil {
		t.Error(err)
	}
	if cert.Certificate == nil {
		t.Error("getClientCertificate: want non-nil Certificate, got nil")
	}
	if cert.PrivateKey == nil {
		t.Error("getClientCertificate: want non-nil PrivateKey, got nil")
	}
}

// This test launches a mock signer binary "test_signer.go" that uses an invalid pem file.
func TestEnterpriseCertificateProxySource_InitializationFailure(t *testing.T) {
	_, err := NewEnterpriseCertificateProxySource("testdata/enterprise_certificate_config_invalid_pem.json")
	if err == nil {
		t.Fatal("Expecting error.")
	}
	if got, want := err.Error(), "failed to retrieve certificate chain: unexpected EOF"; got != want {
		t.Errorf("NewEnterpriseCertificateProxySource: want err %v, got %v", want, got)
	}
}
