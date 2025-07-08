package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Zero(t *testing.T) {
	v := EncodeInt(0)
	assert.Equal(t, "", v)
}

func Test_Positive(t *testing.T) {
	for num := 1; num < 10; num++ {
		encoded := EncodeInt(num)
		decoded, err := DecodeInt(encoded)

		assert.Equal(t, num, decoded)
		assert.Nil(t, err)
	}
}

func Test_Error_DecodeBadHash(t *testing.T) {
	num, err := DecodeInt("abcd`")

	assert.Equal(t, 0, num)
	assert.NotNil(t, err)
}
