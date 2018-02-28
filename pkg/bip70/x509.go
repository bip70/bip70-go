package bip70

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

var (
	// ErrUnsupportedKeyType is returned when we really have no
	// idea what kind of key is involved.
	ErrUnsupportedKeyType = errors.New("unsupported key type")

	// ErrUnsupportedPkiType is returned when we don't recognize
	// the pki type in the request, or in input.
	ErrUnsupportedPkiType = errors.New("unsupported pki type")
)

// GetSignatureAlgorithm takes a pkiType, and a certificate,
// extracts the subject public keys type, and produces
// a signature algorithm and a hash function, or an error
func GetSignatureAlgorithm(pkiType string, x509Cert *x509.Certificate) (
	x509.SignatureAlgorithm, crypto.Hash, error) {
	if _, ok := x509Cert.PublicKey.(*rsa.PublicKey); ok {
		if pkiType == PkiTypeX509Sha1 {
			return x509.SHA1WithRSA, crypto.SHA1, nil
		} else if pkiType == PkiTypeX509Sha256 {
			return x509.SHA256WithRSA, crypto.SHA256, nil
		} else {
			return 0, 0, ErrUnsupportedPkiType
		}
	} else if _, ok := x509Cert.PublicKey.(*rsa.PublicKey); ok {
		if pkiType == PkiTypeX509Sha1 {
			return x509.ECDSAWithSHA1, crypto.SHA1, nil
		} else if pkiType == PkiTypeX509Sha256 {
			return x509.ECDSAWithSHA256, crypto.SHA256, nil
		} else {
			return 0, 0, ErrUnsupportedPkiType
		}
	}

	return 0, 0, ErrUnsupportedKeyType
}
