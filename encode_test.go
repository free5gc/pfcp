package pfcp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/pfcp/pfcpType"
)

func TestMarshal(t *testing.T) {
	testCases := []struct {
		name        string
		message     Message
		expectedBuf []byte
	}{
		{
			name: "Association Setup Request",
			message: Message{
				Header: Header{
					Version:         1,
					MP:              0,
					S:               0,
					MessageType:     PFCP_ASSOCIATION_SETUP_REQUEST,
					MessageLength:   0,
					SEID:            0,
					SequenceNumber:  1,
					MessagePriority: 0,
				},
				Body: PFCPAssociationSetupRequest{
					NodeID: &pfcpType.NodeID{
						NodeIdType: 0,
						IP:         net.ParseIP("192.188.2.2").To4(),
					},
				},
			},
			expectedBuf: []byte{0x20, 0x5, 0x0, 0xd, 0x0, 0x0, 0x1, 0x0, 0x0, 0x3c, 0x0, 0x5, 0x0, 0xc0, 0xbc, 0x2, 0x2},
		},
		{
			name: "Session Establishment Request",
			message: Message{
				Header: Header{
					Version:         1,
					MP:              0,
					S:               1,
					MessageType:     PFCP_SESSION_ESTABLISHMENT_REQUEST,
					MessageLength:   0,
					SEID:            0x02,
					SequenceNumber:  1,
					MessagePriority: 0,
				},
				Body: PFCPSessionEstablishmentRequest{
					NodeID: &pfcpType.NodeID{
						NodeIdType: 0,
						IP:         net.ParseIP("192.188.2.2").To4(),
					},
					CreatePDR: []*CreatePDR{
						{
							PDRID:              &pfcpType.PacketDetectionRuleID{RuleId: 1},
							Precedence:         &pfcpType.Precedence{PrecedenceValue: 32},
							OuterHeaderRemoval: &pfcpType.OuterHeaderRemoval{OuterHeaderRemovalDescription: pfcpType.OuterHeaderRemovalGtpUUdpIpv4},
						},
					},
				},
			},
			expectedBuf: []byte{0x21, 0x32, 0x0, 0x2c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x1, 0x0, 0x0, 0x3c, 0x0, 0x5, 0x0, 0xc0, 0xbc, 0x2, 0x2, 0x0, 0x1, 0x0, 0x13, 0x0, 0x38, 0x0, 0x2, 0x0, 0x1, 0x0, 0x1d, 0x0, 0x4, 0x0, 0x0, 0x0, 0x20, 0x0, 0x5f, 0x0, 0x1, 0x0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf, err := tc.message.Marshal()
			require.Nil(t, err, "encode message", tc.name, "failed")
			require.Equal(t, tc.expectedBuf, buf)
		})
	}
}
