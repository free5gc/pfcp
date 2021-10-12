package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQFIMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name string
		qfi  QFI
		buf  []byte
	}{
		{
			name: "QFI 1",
			qfi: QFI{
				QFI: 1,
			},
			buf: []byte{0x01},
		},
		{
			name: "QFI 63",
			qfi: QFI{
				QFI: 63,
			},
			buf: []byte{0x3F},
		},
		{
			name: "QFI 17",
			qfi: QFI{
				QFI: 17,
			},
			buf: []byte{0x11},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.qfi.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var qfi QFI
				err := qfi.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.qfi, qfi)
			})
		}
	})
}
