// Copyright 2022 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// test_signer.go is a net/rpc server that listens on stdin/stdout, exposing
// mock methods for testing enterprise certificate proxy flow.
package main

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"time"
)

// SignArgs encapsulate the parameters for the Sign method.
type SignArgs struct {
	Digest []byte
	Opts   crypto.SignerOpts
}

// EnterpriseCertSigner exports RPC methods for signing.
type EnterpriseCertSigner struct {
	cert *tls.Certificate
}

// Transport wraps a pair of unidirectional streams as an io.ReadWriteCloser.
type Transport struct {
	io.ReadCloser
	io.WriteCloser
}

// Close closes t's underlying ReadCloser and WriteCloser.
func (t *Transport) Close() error {
	rerr := t.ReadCloser.Close()
	werr := t.WriteCloser.Close()
	if rerr != nil {
		return rerr
	}
	return werr
}

// CertificateChain returns the credential as a raw X509 cert chain. This
// contains the public key.
func (k *EnterpriseCertSigner) CertificateChain(ignored struct{}, certificateChain *[][]byte) error {
	*certificateChain = k.cert.Certificate
	return nil
}

// Public returns the first public key for this Key, in ASN.1 DER form.
func (k *EnterpriseCertSigner) Public(ignored struct{}, publicKey *[]byte) (err error) {
	if len(k.cert.Certificate) == 0 {
		return nil
	}
	cert, err := x509.ParseCertificate(k.cert.Certificate[0])
	*publicKey, err = x509.MarshalPKIXPublicKey(cert.PublicKey)
	return nil
}

// Sign signs a message by encrypting a message digest.
func (k *EnterpriseCertSigner) Sign(args SignArgs, resp *[]byte) (err error) {
	return nil
}

func main() {
	enterpriseCertSigner := new(EnterpriseCertSigner)

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading certificate: %v", err)
	}
	cert, err := tls.X509KeyPair(data, data)
	if err != nil {
		log.Fatalf("Error creating X509 certificate: %v", err)
	}

	enterpriseCertSigner.cert = &cert

	if err := rpc.Register(enterpriseCertSigner); err != nil {
		log.Fatalf("Error registering net/rpc: %v", err)
	}

	// If the parent process dies, we should exit.
	// We can detect this by periodically checking if the PID of the parent
	// process is 1 (https://stackoverflow.com/a/2035683).
	go func() {
		for {
			if os.Getppid() == 1 {
				log.Fatalln("Parent process died, exiting...")
			}
			time.Sleep(time.Second)
		}
	}()

	rpc.ServeConn(&Transport{os.Stdin, os.Stdout})
}
