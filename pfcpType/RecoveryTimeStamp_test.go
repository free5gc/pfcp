package pfcpType

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalRecoveryTimeStamp(t *testing.T) {
	testData := RecoveryTimeStamp{
		RecoveryTimeStamp: time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{135, 108, 229, 128}, buf)
}

func TestUnmarshalRecoveryTimeStamp(t *testing.T) {
	buf := []byte{135, 108, 229, 128}
	var testData RecoveryTimeStamp
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := RecoveryTimeStamp{
		RecoveryTimeStamp: time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	assert.Equal(t, expectData, testData)
}
