package bip70

import (
	_assert "github.com/stretchr/testify/require"
	"testing"
)

func TestConst(t *testing.T) {
	assert := _assert.New(t)
	assert.Equal("none", PkiTypeNone)
	assert.Equal("x509+sha1", PkiTypeX509Sha1)
	assert.Equal("x509+sha256", PkiTypeX509Sha256)
}
