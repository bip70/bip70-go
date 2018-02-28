package bip70

import (
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"github.com/golang/protobuf/proto"
)

// Builder is an interface designed for completing
// a payment request. It receives PaymentDetails, and
// renders it into a final PaymentRequest object.
type Builder interface {
	Build(details *payments.PaymentDetails) (*payments.PaymentRequest, error)
}

// NewUnsignedBuilder returns an UnsignedBuilder.
func NewUnsignedBuilder() (*UnsignedBuilder, error) {
	return &UnsignedBuilder{}, nil
}

// UnsignedBuilder implements the Builder interface,
// and works by embedding the payment details into
// a request, without a signature (PkiTypeNone)
type UnsignedBuilder struct {
}

// Build takes the details struct and converts it to
// a payment request, or returns an error on failure.
func (c *UnsignedBuilder) Build(details *payments.PaymentDetails) (*payments.PaymentRequest, error) {

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
