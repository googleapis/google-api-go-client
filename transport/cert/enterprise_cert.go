// Copyright 2022 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cert

import (
	"crypto/tls"
	"errors"
	"os"

	"github.com/googleapis/enterprise-certificate-proxy/client"
)

type ecpSource struct {
	key *client.Key
}

// NewEnterpriseCertificateProxySource creates a certificate source
// using the enterprise-certificate-proxy client, which delegates
// certifcate related operations to an OS-specific "signer binary"
// that communicates with the native keystore (ex. keychain on MacOS).
//
// The configFilePath points to a config file containing relevant parameters
// such as the certificate issuer and the location of the signer binary.
// If configFilePath is empty, the client will attempt to load the config from
// a well-known gcloud location.
//
// Return nil for Source and Error if config file is missing.
func NewEnterpriseCertificateProxySource(configFilePath string) (Source, error) {
	key, err := client.Cred(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		// Ignore.
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return (&ecpSource{
		key: key,
	}).getClientCertificate, nil
}

func (s *ecpSource) getClientCertificate(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
	var cert tls.Certificate
	cert.PrivateKey = s.key
	cert.Certificate = s.key.CertificateChain()
	return &cert, nil
}
