package pfcp

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NodeRelatedHeaderHex    = "2005000000000100"
	SessionRelatedHeaderHex = "233400000000000000000001000000F0"
)

var NodeRelatedHeader = Header{
	Version:        1,
	S:              0,
	MessageType:    PFCP_ASSOCIATION_SETUP_REQUEST,
	MessageLength:  0,
	SequenceNumber: 1,
}

var SessionRelatedHeader = Header{
	Version:         1,
	MP:              1,
	S:               1,
	MessageType:     PFCP_SESSION_MODIFICATION_REQUEST,
	MessageLength:   0,
	SEID:            1,
	SequenceNumber:  0,
	MessagePriority: 15,
}

func TestPFCPHeader_MarshalBinary(t *testing.T) {
	var payload []byte
	var expect []byte

	// Test PFCP Header for Node Related Messages
	payload, _ = NodeRelatedHeader.MarshalBinary()
	expect, _ = hex.DecodeString(NodeRelatedHeaderHex)
	assert.Equal(t, expect, payload)

	// Test PFCP Header for Session Related Messages
	payload, _ = SessionRelatedHeader.MarshalBinary()
	expect, _ = hex.DecodeString(SessionRelatedHeaderHex)
	assert.Equal(t, expect, payload)
}

func TestPFCPHeader_UnmarshalBinary(t *testing.T) {
	var tmpHeader Header
	var payload []byte

	// Test PFCP Header for Node Related Messages
	payload, _ = hex.DecodeString(NodeRelatedHeaderHex)
	_ = tmpHeader.UnmarshalBinary(payload)
	assert.Equal(t, NodeRelatedHeader, tmpHeader)

	// Test PFCP Header for Session Related Messages
	payload, _ = hex.DecodeString(SessionRelatedHeaderHex)
	_ = tmpHeader.UnmarshalBinary(payload)
	assert.Equal(t, SessionRelatedHeader, tmpHeader)
}

func TestPFCPHeaderUnmarshalBinaryRejectsTruncatedSessionHeader(t *testing.T) {
	testCases := []struct {
		name      string
		input     []byte
		errorText string
	}{
		{
			name:      "missing SEID",
			input:     []byte{0x21, byte(PFCP_SESSION_ESTABLISHMENT_REQUEST), 0x00, 0x04, 0x00, 0x00, 0x01, 0x00},
			errorText: "failed to read PFCP SEID",
		},
		{
			name:      "missing sequence number",
			input:     []byte{0x21, byte(PFCP_SESSION_ESTABLISHMENT_REQUEST), 0x00, 0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
			errorText: "failed to read PFCP sequence number",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var header Header
			err := header.UnmarshalBinary(tc.input)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.errorText)
		})
	}
}
