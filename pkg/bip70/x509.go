package bip70

import (
	"crypto/x509"
	"errors"
	"crypto/rsa"
	"crypto"
)

var (
	// ErrNoCertificates is returned when certificate chain
	// validation is attempted but with no certificates
	ErrNoCertificates = errors.New("no certificates in bundle")
)

func GetSignatureAlgorithm(pkiType string, x509Cert *x509.Certificate) (
	x509.SignatureAlgorithm, crypto.Hash, error) {
	if _, ok := x509Cert.PublicKey.(*rsa.PublicKey); ok {
		if pkiType == PkiTypeX509Sha1 {
			return x509.SHA1WithRSA, crypto.SHA1, nil
		} else if pkiType == PkiTypeX509Sha256 {
			return x509.SHA256WithRSA, crypto.SHA256, nil
		}
	} else if _, ok := x509Cert.PublicKey.(*rsa.PublicKey); ok {
		if pkiType == PkiTypeX509Sha1 {
			return x509.ECDSAWithSHA1, crypto.SHA1, nil
		} else if pkiType == PkiTypeX509Sha256 {
			return x509.ECDSAWithSHA256, crypto.SHA256, nil
		}
	}

	return 0, 0, errors.New("unsupported public key type")
}
