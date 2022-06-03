// Copyright 2022 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cert

import (
	"bytes"
	"testing"
)

func TestSecureConnectSource_ConfigMissing(t *testing.T) {
	source, err := NewSecureConnectSource("missing.json")
	if err != nil {
		t.Fatal("NewSecureConnectSource: expected nil error returned when config is missing.")
	}
	if source != nil {
		t.Error("NewSecureConnectSource: expected nil source and error returned when config is missing.")
	}
}

func TestSecureConnectSource_GetClientCertificateSuccess(t *testing.T) {
	source, err := NewSecureConnectSource("testdata/context_aware_metadata.json")
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

func TestSecureConnectSource_GetClientCertificateFailure(t *testing.T) {
	source, err := NewSecureConnectSource("testdata/context_aware_metadata_invalid_pem.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = source(nil)
	if err == nil {
		t.Error("Expecting error.")
	}
	if got, want := err.Error(), "x509: malformed certificate"; got != want {
		t.Errorf("getClientCertificate: want %v err, got %v", want, got)
	}
}

func TestSecureConnectSource_ValidateMetadataSuccess(t *testing.T) {
	metadata := secureConnectMetadata{Cmd: []string{"cat", "testdata/testcert.pem"}}
	err := validateMetadata(metadata)
	if err != nil {
		t.Error(err)
	}
}

func TestSecureConnectSource_ValidateMetadataFailure(t *testing.T) {
	metadata := secureConnectMetadata{Cmd: []string{}}
	err := validateMetadata(metadata)
	if err == nil {
		t.Error("validateMetadata: want non-nil err, got nil")
	}
	if want, got := "empty cert_provider_command", err.Error(); want != got {
		t.Errorf("validateMetadata: want %v err, got %v", want, got)
	}
}

func TestSecureConnectSource_IsCertificateExpiredTrue(t *testing.T) {
	source, err := NewSecureConnectSource("testdata/context_aware_metadata.json")
	if err != nil {
		t.Fatal(err)
	}
	cert, err := source(nil)
	if err != nil {
		t.Error(err)
	}
	if !isCertificateExpired(cert) {
		t.Error("isCertificateExpired: want true, got false")
	}
}

func TestSecureConnectSource_IsCertificateExpiredFalse(t *testing.T) {
	source, err := NewSecureConnectSource("testdata/context_aware_metadata_nonexpiring_pem.json")
	if err != nil {
		t.Fatal(err)
	}
	cert, err := source(nil)
	if err != nil {
		t.Error(err)
	}
	if isCertificateExpired(cert) {
		t.Error("isCertificateExpired: want false, got true")
	}
}

func TestCertificateCaching(t *testing.T) {
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/nonexpiring.pem"}}}
	cert, err := source.getClientCertificate(nil)
	if err != nil {
		t.Fatal(err)
	}
	if cert == nil {
		t.Error("getClientCertificate: want non-nil cert, got nil")
	}
	if source.cachedCert == nil {
		t.Error("getClientCertificate: want non-nil cachedCert, got nil")
	}
	if !bytes.Equal(cert.Certificate[0], source.cachedCert.Certificate[0]) {
		t.Error("Cached certificate is different.")
	}
	if cert.PrivateKey != source.cachedCert.PrivateKey {
		t.Error("Cached PrivateKey is different.")
	}
}
