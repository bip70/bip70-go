package bip70

import (
	"crypto/x509"
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
)

// NewX509Certificates takes an entityCertificate, a slice
// of intermediate certificates, and produces an X509Certificates
// protobuf message
func NewX509Certificates(cert *x509.Certificate, intermediates []*x509.Certificate) *payments.X509Certificates {
	certs := &payments.X509Certificates{
		Certificate: make([][]byte, 0, len(intermediates)+1),
	}

	certs.Certificate[0] = cert.Raw
	for i := 0; i < len(intermediates); i++ {
		certs.Certificate[i+1] = intermediates[i].Raw
	}

	return certs
}

// ParseX509Certificates extracts a X509Certificates message
// into an entity certificate, and a CertPool containing
// the intermediate certificates.
func ParseX509Certificates(certs *payments.X509Certificates) (*x509.Certificate, *x509.CertPool, error) {
	nIn := len(certs.Certificate)
	if nIn < 1 {
		return nil, nil, ErrNoCertificates
	}

	entityCert, err := x509.ParseCertificate(certs.Certificate[0])
	if err != nil {
		return nil, nil, err
	}

	intermediates := x509.NewCertPool()
	for i := 1; i < len(certs.Certificate); i++ {
		intermediate, err := x509.ParseCertificate(certs.Certificate[i])
		if err != nil {
			return nil, nil, err
		}
		intermediates.AddCert(intermediate)
	}

	return entityCert, intermediates, nil
}
