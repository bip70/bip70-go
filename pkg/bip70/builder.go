package bip70

import (
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"github.com/golang/protobuf/proto"
)

type Builder interface {
	Build(details *payments.PaymentDetails) (*payments.PaymentRequest, error)
}

func NewUnsignedBuilder() (*UnsignedBuilder, error) {
	return &UnsignedBuilder{}, nil
}

type UnsignedBuilder struct {
}

func (c *UnsignedBuilder) Build(details *payments.PaymentDetails) (
	*payments.PaymentRequest, error) {

	detailsBin, err := proto.Marshal(details)
	if err != nil {
		return nil, err
	}

	version := uint32(1)
	pkiType := PkiTypeNone
	req := &payments.PaymentRequest{}
	req.PaymentDetailsVersion = &version
	req.SerializedPaymentDetails = detailsBin
	req.PkiType = &pkiType

	return req, nil
}
