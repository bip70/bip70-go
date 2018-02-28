package bip70

import (
	"time"
	"crypto/x509"
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"errors"
	"github.com/golang/protobuf/proto"
)

type ValidationConfig struct {
	DnsName   string
	Time      time.Time
	RootPool  *x509.CertPool
	KeyUsages []x509.ExtKeyUsage
}

type CheckerInterface interface {
	ValidateChain(cfg *ValidationConfig, derCerts *payments.X509Certificates) (
		[][]*x509.Certificate, error)
	ValidateSignature(cert *x509.Certificate, req *payments.PaymentRequest) (
	error)
}

type X509Checker struct {
}

func (c *X509Checker) ValidateChain(cfg *ValidationConfig,
	certs *payments.X509Certificates) ([][]*x509.Certificate, error) {
	nIn := len(certs.Certificate)
	if nIn < 1 {
		return nil, ErrNoCertificates
	}

	entityCert, err := x509.ParseCertificate(certs.Certificate[0])
	if err != nil {
		return nil, err
	}

	intermediates := x509.NewCertPool()
	for i := 1; i < nIn; i++ {
		intermediate, err := x509.ParseCertificate(certs.Certificate[i])
		if err != nil {
			return nil, err
		}
		intermediates.AddCert(intermediate)
	}

	chains, err := entityCert.Verify(x509.VerifyOptions{
		DNSName: cfg.DnsName,
		Roots: cfg.RootPool,
		Intermediates: intermediates,
		CurrentTime: cfg.Time,
		KeyUsages: cfg.KeyUsages,
	})
	if err != nil {
		return nil, err
	}

	return chains, nil
}
func (c *X509Checker) ValidateSignature(cert *x509.Certificate,
	req *payments.PaymentRequest) error {
	if len(req.GetSignature()) < 1 {
		return errors.New("empty signature")
	}

	sigAlg, _, err := GetSignatureAlgorithm(req.GetPkiType(), cert)
	if err != nil {
		return err
	}

	clone := proto.Clone(req)
	reqCpy, ok := clone.(*payments.PaymentRequest)
	if !ok {
		return errors.New("failed to clone request")
	}
	reqCpy.Signature = []byte{}

	signData, err := proto.Marshal(reqCpy)
	if err != nil {
		return err
	}

	return cert.CheckSignature(sigAlg, signData, req.Signature)
}
