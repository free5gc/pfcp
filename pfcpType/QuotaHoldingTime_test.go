package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalQuotaHoldingTime(t *testing.T) {
	testData := QuotaHoldingTime{
		QuotaHoldingTimeValue: 12345,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{0, 0, 48, 57}, buf)
}

func TestUnmarshalQuotaHoldingTime(t *testing.T) {
	buf := []byte{0, 0, 48, 57}
	var testData QuotaHoldingTime
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := QuotaHoldingTime{
		QuotaHoldingTimeValue: 12345,
	}
	assert.Equal(t, expectData, testData)
}
