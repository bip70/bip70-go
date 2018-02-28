package bip70

import (
	"crypto/x509"
	"errors"
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"github.com/golang/protobuf/proto"
	"time"
)

var (
	// ErrNoCertificates is returned when certificate chain
	// validation is attempted but with no certificates
	ErrNoCertificates = errors.New("no certificates in bundle")

	// ErrEmptySignature is returned when the signature field
	// is empty
	ErrEmptySignature = errors.New("empty signature")
)

// ValidationConfig captures high level validation
// configuration. If RootPool is empty, the system
// store is used as a default. See x509.VerifyOptions
// for additional information, this struct exists to
// supplement details not known by parsing X509 certs,
// to constrain validation, or to unit test (current time)
type ValidationConfig struct {
	DNSName   string
	Time      time.Time
	RootPool  *x509.CertPool
	KeyUsages []x509.ExtKeyUsage
}

// CheckerInterface is a type capable of validating
// a certificate chain, and validate x509 signatures
type CheckerInterface interface {
	// ValidateChain takes a validation config, and a set of
	// certificates, and returns a set of valid chains to
	// the entity certificate, or an error.
	ValidateChain(cfg *ValidationConfig, derCerts *payments.X509Certificates) ([][]*x509.Certificate, error)

	// ValidateSignature takes a certificate, and a
	// payment request, and validates the signature.
	// An error is returned if the signature is invalid.
	ValidateSignature(cert *x509.Certificate, req *payments.PaymentRequest) error
}

// X509Checker implements the CheckerInterface
// performing actual validation on inputs.
type X509Checker struct {
}

// ValidateChain - see CheckerInterface.ValidateChain
func (c *X509Checker) ValidateChain(cfg *ValidationConfig, certs *payments.X509Certificates) ([][]*x509.Certificate, error) {

	entityCert, intermediates, err := ParseX509Certificates(certs)
	if err != nil {
		return nil, err
	}

	chains, err := entityCert.Verify(x509.VerifyOptions{
		DNSName:       cfg.DNSName,
		Roots:         cfg.RootPool,
		Intermediates: intermediates,
		CurrentTime:   cfg.Time,
		KeyUsages:     cfg.KeyUsages,
	})
	if err != nil {
		return nil, err
	}

	return chains, nil
}

// ValidateSignature - see CheckerInterface.ValidateSignature
func (c *X509Checker) ValidateSignature(cert *x509.Certificate, req *payments.PaymentRequest) error {
	if len(req.GetSignature()) < 1 {
		return ErrEmptySignature
	}

	sigAlg, _, err := GetSignatureAlgorithm(req.GetPkiType(), cert)
	if err != nil {
		return err
	}

	reqCpy := proto.Clone(req).(*payments.PaymentRequest)
	reqCpy.Signature = []byte{}

	signData, err := proto.Marshal(reqCpy)
	if err != nil {
		return err
	}

	return cert.CheckSignature(sigAlg, signData, req.Signature)
}
