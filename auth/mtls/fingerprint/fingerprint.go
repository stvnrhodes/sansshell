package fingerprint

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

func WithSha256Check(config *tls.Config, fingerprint []byte) *tls.Config {
	fingerprintCreds := config.Clone()
	// Since this tends to get used with invalid certs, we must set
	// InsecureSkipVerify to skip the normal TLS checks.
	// VerifyConnection will still run when InsecureSkipVerify
	// is true.
	fingerprintCreds.InsecureSkipVerify = true
	fingerprintCreds.VerifyConnection = func(cs tls.ConnectionState) error {
		hash := Sha256(cs.PeerCertificates[0])
		if !bytes.Equal(hash, fingerprint) {
			return fmt.Errorf("unexpected cert fingerprint %v", hash)
		}
		return nil
	}
	return fingerprintCreds
}

func Sha256(cert *x509.Certificate) []byte {
	hash := sha256.Sum256(cert.Raw)
	return hash[:]
}
