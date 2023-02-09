package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalMeasurementInformation(t *testing.T) {
	testData := MeasurementInformation{
		Mnop: true,
		Istm: false,
		Radi: true,
		Inam: false,
		Mbqe: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{21}, buf)
}

func TestUnmarshalMeasurementInformation(t *testing.T) {
	buf := []byte{21}
	var testData MeasurementInformation
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := MeasurementInformation{
		Mnop: true,
		Istm: false,
		Radi: true,
		Inam: false,
		Mbqe: true,
	}
	assert.Equal(t, expectData, testData)
}
