package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQERIDMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name  string
		qerID QERID
		buf   []byte
	}{
		{
			name: "QERID 1",
			qerID: QERID{
				QERID: 1,
			},
			buf: []byte{0x00, 0x00, 0x00, 0x01},
		},
		{
			name: "QERID 1",
			qerID: QERID{
				QERID: 256,
			},
			buf: []byte{0x00, 0x00, 0x01, 0x00},
		},
		{
			name: "QERID 458715885",
			qerID: QERID{
				QERID: 458715885,
			},
			buf: []byte{0x1B, 0x57, 0x72, 0xED},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.qerID.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var qerID QERID
				err := qerID.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.qerID, qerID)
			})
		}
	})
}
