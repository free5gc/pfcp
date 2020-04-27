package pfcp_test

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"free5gc/lib/pfcp"
	"free5gc/lib/pfcp/pfcpType"
	"net"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	associationHex := "2005000d00000100003c000500c0bc0202"
	b, _ := hex.DecodeString(associationHex)
	message := pfcp.Message{}
	err := message.Unmarshal(b)
	assert.Equal(t, pfcp.Message{
		Header: pfcp.Header{
			Version:         1,
			MP:              0,
			S:               0,
			MessageType:     pfcp.PFCP_ASSOCIATION_SETUP_REQUEST,
			MessageLength:   13,
			SEID:            0,
			SequenceNumber:  1,
			MessagePriority: 0,
		},
		Body: pfcp.PFCPAssociationSetupRequest{
			NodeID: &pfcpType.NodeID{
				NodeIdType:  0,
				NodeIdValue: net.ParseIP("192.188.2.2").To4(),
			},
		},
	},
		message)
	assert.Nil(t, err)
}
