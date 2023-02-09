package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalMultiplier(t *testing.T) {
	testData := Multiplier{
		Multiplierdata: []byte("test"),
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{116, 101, 115, 116}, buf)
}

func TestUnmarshalMultiplier(t *testing.T) {
	buf := []byte{116, 101, 115, 116}
	var testData Multiplier
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := Multiplier{
		Multiplierdata: []byte("test"),
	}
	assert.Equal(t, expectData, testData)
}
