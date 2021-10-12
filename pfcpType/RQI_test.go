package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRQIMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name string
		rqi  RQI
		buf  []byte
	}{
		{
			name: "RQI on",
			rqi: RQI{
				RQI: true,
			},
			buf: []byte{0x01},
		},
		{
			name: "RQI off",
			rqi: RQI{
				RQI: false,
			},
			buf: []byte{0x00},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.rqi.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var rqi RQI
				err := rqi.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.rqi, rqi)
			})
		}
	})
}
