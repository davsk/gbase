// /////////////////////////////////////////////////////////////
// 'Cert.go'                                                   /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 25, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/goinggo/tracelog"
)

type Cert struct {
	HostName string
	CertFile string
	KeyFile  string
}

func (ct *Cert) Default(name string) {
	ct.HostName = name
	ct.CertFile = "Cert.pem"
	ct.KeyFile = "key.pem"
}

func (ct *Cert) Validate() {
	// const FunctionName for tracelog.
	const traceFunctionName = "Cert.Validate"
	tracelog.Started(traceTitle, traceFunctionName)

	_ = ct

	if _, err := os.Stat(ct.CertFile); os.IsExist(err) {
		if _, err := os.Stat(ct.KeyFile); os.IsExist(err) {
			// TODO(dls) check Cert date for expiration.
			// Completed successfully.
			tracelog.Completed(traceTitle, traceFunctionName)

			return
		}
	}

	ct.Generate()

	// Completed successfully.
	tracelog.Completed(traceTitle, traceFunctionName)
}

func (ct *Cert) Generate() (err error) {
	// const FunctionName for tracelog.
	const traceFunctionName = "Cert.Generate"
	tracelog.Started(traceTitle, traceFunctionName)

	var priv interface{}

	// Generate Key
	priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return err
	}
	tracelog.Info(traceTitle, traceFunctionName,
		"Private Key := %v", priv)

	// Temporal calculations
	notBefore := time.Now()
	validFor := 365 * 24 * time.Hour
	notAfter := notBefore.Add(validFor)
	tracelog.Info(traceTitle, traceFunctionName,
		"notAfter := %v", notAfter)

	// Big Random Number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		tracelog.CompletedError(err, traceTitle, traceFunctionName)
		return err
	}
	tracelog.Info(traceTitle, traceFunctionName,
		"serialNumber := %v", serialNumber)

	// Template for Certificate
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Davsk℠ Universe 4.0 Game"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Parse HostName
	hosts := strings.Split(ct.HostName, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	// if isCA
	if true {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	// Create certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		return err
	}

	certOut, err := os.Create(ct.CertFile)
	if err != nil {
		return err
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	tracelog.Info(traceTitle, traceFunctionName,
		"CertFile := %v written.", ct.CertFile)

	// Create Key File
	keyOut, err := os.OpenFile(ct.KeyFile, os.O_WRONLY|os.O_CREATE|os.
		O_TRUNC, 0600)
	if err != nil {
		log.Print("failed to open key.pem for writing:", err)
		return
	}
	pem.Encode(keyOut, pemBlockForKey(priv))
	keyOut.Close()
	tracelog.Info(traceTitle, traceFunctionName,
		"KeyFile := %v wriiten.", ct.KeyFile)

	return nil
}

// derived from https://golang.org/src/crypto/tls/generate_cert.go
func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// derived from https://golang.org/src/crypto/tls/generate_cert.go
func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}
