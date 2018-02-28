package bip70_test

import (
	"bytes"
	"github.com/bip70/bip70-go/pkg/bip70"
	"github.com/bip70/bip70-go/pkg/protobuf/payments"
	"github.com/golang/protobuf/proto"
	_assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewUnsignedBuilder(t *testing.T) {
	assert := _assert.New(t)
	builder, err := bip70.NewUnsignedBuilder()
	assert.Nil(err)
	assert.NotNil(builder)
}

func TestBuildInvalidDetails(t *testing.T) {
	assert := _assert.New(t)
	builder, err := bip70.NewUnsignedBuilder()
	assert.Nil(err)
	assert.NotNil(builder)

	det := &payments.PaymentDetails{}
	req, err := builder.Build(det)
	assert.Error(err)
	assert.EqualError(err, "proto: required field \"Time\" not set")
	assert.Nil(req)
}

func TestBuildUnsigned(t *testing.T) {
	assert := _assert.New(t)
	builder, err := bip70.NewUnsignedBuilder()
	assert.Nil(err)
	assert.NotNil(builder)

	now := uint64(time.Now().Second())
	det := &payments.PaymentDetails{
		Time: &now,
	}
	req, err := builder.Build(det)
	assert.Nil(err)
	assert.NotNil(req)

	assert.Equal(bip70.PkiTypeNone, req.GetPkiType())
	assert.Equal(uint32(1), req.GetPaymentDetailsVersion())
	assert.Nil(req.GetPkiData())
	assert.Nil(req.GetSignature())

	deSer := &payments.PaymentDetails{}
	err = proto.Unmarshal(req.SerializedPaymentDetails, deSer)
	assert.NoError(err)

	ser, err := proto.Marshal(deSer)
	assert.NoError(err)
	assert.True(bytes.Equal(ser, req.SerializedPaymentDetails))
}
