// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cert

import (
	"bytes"
	"testing"
)

func TestGetClientCertificateSuccess(t *testing.T) {
	defaultCert.cachedCert = nil
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/testcert.pem"}}}
	cert, err := source.getClientCertificate(nil)
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

func TestGetClientCertificateFailure(t *testing.T) {
	defaultCert.cachedCert = nil
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat"}}}
	_, err := source.getClientCertificate(nil)
	if err == nil {
		t.Error("Expecting error.")
	}
	if got, want := err.Error(), "tls: failed to find any PEM data in certificate input"; got != want {
		t.Errorf("getClientCertificate: want %v err, got %v", want, got)
	}
}

func TestValidateMetadataSuccess(t *testing.T) {
	metadata := secureConnectMetadata{Cmd: []string{"cat", "testdata/testcert.pem"}}
	err := validateMetadata(metadata)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateMetadataFailure(t *testing.T) {
	metadata := secureConnectMetadata{Cmd: []string{}}
	err := validateMetadata(metadata)
	if err == nil {
		t.Error("validateMetadata: want non-nil err, got nil")
	}
	if want, got := "empty cert_provider_command", err.Error(); want != got {
		t.Errorf("validateMetadata: want %v err, got %v", want, got)
	}
}

func TestIsCertificateExpiredTrue(t *testing.T) {
	defaultCert.cachedCert = nil
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/testcert.pem"}}}
	cert, err := source.getClientCertificate(nil)
	if err != nil {
		t.Error(err)
	}
	if !isCertificateExpired(cert) {
		t.Error("isCertificateExpired: want true, got false")
	}
}

func TestIsCertificateExpiredFalse(t *testing.T) {
	defaultCert.cachedCert = nil
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/nonexpiringtestcert.pem"}}}
	cert, err := source.getClientCertificate(nil)
	if err != nil {
		t.Error(err)
	}
	if isCertificateExpired(cert) {
		t.Error("isCertificateExpired: want false, got true")
	}
}

func TestCertificateCaching(t *testing.T) {
	defaultCert.cachedCert = nil
	source := secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/nonexpiringtestcert.pem"}}}
	cert, err := source.getClientCertificate(nil)
	if err != nil {
		t.Error(err)
	}
	if cert == nil {
		t.Error("getClientCertificate: want non-nil cert, got nil")
	}
	if defaultCert.cachedCert == nil {
		t.Error("getClientCertificate: want non-nil defaultSourceCachedCert, got nil")
	}

	source = secureConnectSource{metadata: secureConnectMetadata{Cmd: []string{"cat", "testdata/testcert.pem"}}}
	cert, err = source.getClientCertificate(nil)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(cert.Certificate[0], defaultCert.cachedCert.Certificate[0]) {
		t.Error("getClientCertificate: want cached Certificate, got different Certificate")
	}
	if cert.PrivateKey != defaultCert.cachedCert.PrivateKey {
		t.Error("getClientCertificate: want cached PrivateKey, got different PrivateKey")
	}
}
