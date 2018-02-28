package bip70

import (
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"crypto/x509"
)

func NewX509Certificates(cert *x509.Certificate, intermediates []*x509.Certificate) (*payments.X509Certificates) {
	certs := &payments.X509Certificates{
		Certificate: make([][]byte, 0, len(intermediates)+1),
	}

	certs.Certificate[0] = cert.Raw
	for i := 0; i < len(intermediates); i++ {
		certs.Certificate[i+1] = intermediates[i].Raw
	}

	return certs
}
