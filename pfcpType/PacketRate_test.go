package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPacketRateMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name       string
		packetRate PacketRate
		buf        []byte
	}{
		{
			name: "UL & DL",
			packetRate: PacketRate{
				ULPR:       true,
				DLPR:       true,
				ULTimeUnit: PacketRateTimeUnitMinute,
				MaximumUL:  1000,
				DLTimeUnit: PacketRateTimeUnitMinute,
				MaximumDL:  40000,
			},
			buf: []byte{0x03, 0x00, 0x03, 0xE8, 0x00, 0x9C, 0x40},
		},
		{
			name: "UL",
			packetRate: PacketRate{
				ULPR:       true,
				DLPR:       false,
				ULTimeUnit: PacketRateTimeUnitMinute,
				MaximumUL:  1000,
			},
			buf: []byte{0x01, 0x00, 0x03, 0xE8},
		},
		{
			name: "DL",
			packetRate: PacketRate{
				ULPR:       false,
				DLPR:       true,
				DLTimeUnit: PacketRateTimeUnitWeek,
				MaximumDL:  40000,
			},
			buf: []byte{0x02, 0x04, 0x9C, 0x40},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.packetRate.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var packetRate PacketRate
				err := packetRate.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.packetRate, packetRate)
			})
		}
	})
}
