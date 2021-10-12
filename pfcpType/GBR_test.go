package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGBRMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name string
		gbr  GBR
		buf  []byte
	}{
		{
			name: "1000&1000",
			gbr: GBR{
				ULGBR: 1000,
				DLGBR: 1000,
			},
			buf: []byte{0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x03, 0xE8},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.gbr.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var gbr GBR
				err := gbr.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.gbr, gbr)
			})
		}
	})
}
