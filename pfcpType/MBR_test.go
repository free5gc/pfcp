package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMBRMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name string
		mbr  MBR
		buf  []byte
	}{
		{
			name: "1000&1000",
			mbr: MBR{
				ULMBR: 1000,
				DLMBR: 1000,
			},
			buf: []byte{0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x03, 0xE8},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.mbr.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var mbr MBR
				err := mbr.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.mbr, mbr)
			})
		}
	})
}
