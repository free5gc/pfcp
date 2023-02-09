package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalVolumeMeasurement(t *testing.T) {
	testData := VolumeMeasurement{
		Dlnop:          true,
		Ulnop:          true,
		Tonop:          true,
		Dlvol:          true,
		Ulvol:          true,
		Tovol:          true,
		TotalVolume:    987654321987654321,
		UplinkVolume:   123456789123456789,
		DownlinkVolume: 864197532864197532,
		TotalPktNum:    135792468135792468,
		UplinkPktNum:   246813579246813579,
		DownlinkPktNum: 765432198765432198,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{63, 13, 180, 218, 95, 126, 244, 18, 177, 1, 182, 155, 75, 172, 208, 95, 21, 11, 254, 63, 19, 210, 35, 179, 156,
		1, 226, 110, 135, 194, 104, 79, 84, 3, 108, 219, 164, 132, 191, 193, 139, 10, 159, 92, 135, 131, 15, 77, 134}, buf)
}

func TestUnmarshalVolumeMeasurement(t *testing.T) {
	buf := []byte{63, 13, 180, 218, 95, 126, 244, 18, 177, 1, 182, 155, 75, 172, 208, 95, 21, 11, 254, 63, 19, 210, 35, 179, 156,
		1, 226, 110, 135, 194, 104, 79, 84, 3, 108, 219, 164, 132, 191, 193, 139, 10, 159, 92, 135, 131, 15, 77, 134}
	var testData VolumeMeasurement
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := VolumeMeasurement{
		Dlnop:          true,
		Ulnop:          true,
		Tonop:          true,
		Dlvol:          true,
		Ulvol:          true,
		Tovol:          true,
		TotalVolume:    987654321987654321,
		UplinkVolume:   123456789123456789,
		DownlinkVolume: 864197532864197532,
		TotalPktNum:    135792468135792468,
		UplinkPktNum:   246813579246813579,
		DownlinkPktNum: 765432198765432198,
	}
	assert.Equal(t, expectData, testData)
}
