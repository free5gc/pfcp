package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUsageReportTrigger(t *testing.T) {
	testData := UsageReportTrigger{
		Immer: true,
		Droth: false,
		Stopt: true,
		Start: false,
		Quhti: true,
		Timth: false,
		Volth: true,
		Perio: false,
		Eveth: true,
		Macar: false,
		Envcl: true,
		Monit: false,
		Termr: true,
		Liusa: false,
		Timqu: true,
		Volqu: false,
		Emrre: true,
		Quvti: false,
		Ipmjl: true,
		Tebur: false,
		Evequ: true,
	}

	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{170, 170, 21}, buf)
}

func TestUnmarshalUsageReportTrigger(t *testing.T) {
	buf := []byte{170, 170, 21}
	var testData UsageReportTrigger
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UsageReportTrigger{
		Immer: true,
		Droth: false,
		Stopt: true,
		Start: false,
		Quhti: true,
		Timth: false,
		Volth: true,
		Perio: false,
		Eveth: true,
		Macar: false,
		Envcl: true,
		Monit: false,
		Termr: true,
		Liusa: false,
		Timqu: true,
		Volqu: false,
		Emrre: true,
		Quvti: false,
		Ipmjl: true,
		Tebur: false,
		Evequ: true,
	}

	assert.Equal(t, expectData, testData)
}
