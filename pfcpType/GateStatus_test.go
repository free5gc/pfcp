package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGateStatusMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name       string
		gateStatus GateStatus
		buf        []byte
	}{
		{
			name: "UL & DL",
			gateStatus: GateStatus{
				ULGate: GateOpen,
				DLGate: GateOpen,
			},
			buf: []byte{0x00},
		},
		{
			name: "UL",
			gateStatus: GateStatus{
				ULGate: GateOpen,
				DLGate: GateClose,
			},
			buf: []byte{0x01},
		},
		{
			name: "DL",
			gateStatus: GateStatus{
				ULGate: GateClose,
				DLGate: GateOpen,
			},
			buf: []byte{0x04},
		},
		{
			name: "None",
			gateStatus: GateStatus{
				ULGate: GateClose,
				DLGate: GateClose,
			},
			buf: []byte{0x05},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := tc.gateStatus.MarshalBinary()
				require.Nil(t, err)
				require.Equal(t, tc.buf, b)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var gateStatus GateStatus
				err := gateStatus.UnmarshalBinary(tc.buf)
				require.Nil(t, err)
				require.Equal(t, tc.gateStatus, gateStatus)
			})
		}
	})
}
