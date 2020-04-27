package pfcp_test

import (
	"fmt"
	"free5gc/lib/pfcp"
	"free5gc/lib/pfcp/pfcpType"
	"net"
	"testing"
)

func TestMarshal(t *testing.T) {
	message := pfcp.Message{
		Header: pfcp.Header{
			Version:         1,
			MP:              0,
			S:               0,
			MessageType:     pfcp.PFCP_ASSOCIATION_SETUP_REQUEST,
			MessageLength:   0,
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
	}
	buf, _ := message.Marshal()
	sip := net.ParseIP("127.0.0.10")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 8805}
	dstAddr := &net.UDPAddr{IP: sip, Port: 8805}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	n, err := conn.Write(buf)
	print(n, err)
	defer conn.Close()
}
