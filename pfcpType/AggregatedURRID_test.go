package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalAggregatedURRID(t *testing.T) {
	testData := AggregatedURRID{
		AggregatedURRIDdata: 12345678,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{0x0, 0xbc, 0x61, 0x4e}, buf)
}

func TestUnmarshalAggregatedURRID(t *testing.T) {
	buf := []byte{0x0, 0xbc, 0x61, 0x4e}
	var testData AggregatedURRID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := AggregatedURRID{
		AggregatedURRIDdata: 12345678,
	}
	assert.Equal(t, expectData, testData)
}
