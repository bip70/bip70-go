package bip70

import (
	"crypto"
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"github.com/golang/protobuf/proto"
	"crypto/rsa"
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/rand"
)

func NewX509Builder(pkiType string, priv crypto.PrivateKey,
	cert *x509.Certificate, intermediates []*x509.Certificate) (
		*X509Builder, error) {

	_, hashFunc, err := GetSignatureAlgorithm(pkiType, cert)
	if err != nil {
		return nil, err
	}

	certs := NewX509Certificates(cert, intermediates)
	certsRaw, err := proto.Marshal(certs)
	if err != nil {
		return nil, err
	}

	return &X509Builder{
		pkiType: pkiType,
		privKey: priv,
		hashFunc: hashFunc,
		cachedCertsBin: certsRaw,
	}, nil
}

type X509Builder struct {
	pkiType string
	privKey crypto.PrivateKey
	hashFunc crypto.SignerOpts
	cachedCertsBin []byte
}

func (c *X509Builder) Build(details *payments.PaymentDetails) (
	*payments.PaymentRequest, error) {

	detailsBin, err := proto.Marshal(details)
	if err != nil {
		return nil, err
	}

	version := uint32(1)
	req := &payments.PaymentRequest{}
	req.PaymentDetailsVersion = &version
	req.SerializedPaymentDetails = detailsBin
	req.PkiType = &c.pkiType
	req.PkiData = c.cachedCertsBin
	req.Signature = []byte{}

	toBeSigned, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	var sig []byte
	switch k := c.privKey.(type) {
	case *rsa.PrivateKey:
		sig, err = k.Sign(rand.Reader, toBeSigned, c.hashFunc)
	case *ecdsa.PrivateKey:
		sig, err = k.Sign(rand.Reader, toBeSigned, c.hashFunc)
	}

	if err != nil {
		return nil, err
	}
	req.Signature = sig
	return req, nil
}
