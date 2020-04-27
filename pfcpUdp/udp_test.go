package pfcpUdp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"free5gc/lib/pfcp"
	"free5gc/lib/pfcp/pfcpType"
)

const testPfcpClientPort = 12345

func TestListen(t *testing.T) {
	testPfcpServerReq := []byte("TestPfcpServerRequest")
	testPfcpServerRsp := []byte("TestPfcpServerResponse")

	server, _ := NewPfcpServer("127.0.0.1")
	defer server.Close()

	go func() {
		buf := make([]byte, MaxPfcpUdpDataSize)
		n, remoteAddr, err := server.Conn.ReadFromUDP(buf)
		assert.Nil(t, err)

		assert.Equal(t, len(testPfcpServerReq), n)
		assert.Equal(t, testPfcpServerReq, buf[:n])
		// t.Log(buf[:n])

		_, err = server.Conn.WriteToUDP(testPfcpServerRsp, remoteAddr)
		assert.Nil(t, err)
	}()

	srcAddr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}
	dstAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: PfcpUdpDestinationPort,
	}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	assert.Nil(t, err)
	defer conn.Close()

	_, err = conn.Write(testPfcpServerReq)
	assert.Nil(t, err)

	buf := make([]byte, MaxPfcpUdpDataSize)
	n, err := conn.Read(buf)
	assert.Nil(t, err)

	assert.Equal(t, len(testPfcpServerRsp), n)
	assert.Equal(t, testPfcpServerRsp, buf[:n])
	// t.Log(buf[:n])
}

func TestPfcpServer(t *testing.T) {
	testPfcpReq := pfcp.Message{
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
				NodeIdValue: net.ParseIP("192.168.1.1").To4(),
			},
		},
	}
	testPfcpRsp := pfcp.Message{
		Header: pfcp.Header{
			Version:         1,
			MP:              0,
			S:               0,
			MessageType:     pfcp.PFCP_ASSOCIATION_SETUP_RESPONSE,
			MessageLength:   13,
			SEID:            0,
			SequenceNumber:  1,
			MessagePriority: 0,
		},
		Body: pfcp.PFCPAssociationSetupResponse{
			NodeID: &pfcpType.NodeID{
				NodeIdType:  0,
				NodeIdValue: net.ParseIP("192.168.1.2").To4(),
			},
		},
	}

	server, _ := NewPfcpServer("127.0.0.1")
	defer server.Close()

	go func() {
		var pfcpMessage pfcp.Message
		remoteAddr, err := server.ReadFrom(&pfcpMessage)
		assert.Nil(t, err)

		assert.Equal(t, testPfcpReq, pfcpMessage)

		err = server.WriteTo(testPfcpRsp, remoteAddr)
		assert.Nil(t, err)
	}()

	srcAddr := &net.UDPAddr{
		IP: net.IPv4zero,
		// IP:   net.ParseIP("127.0.0.1"),
		Port: testPfcpClientPort,
	}
	dstAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: PfcpUdpDestinationPort,
	}

	err := SendPfcpMessage(testPfcpReq, srcAddr, dstAddr)
	assert.Nil(t, err)

	var pfcpMessage pfcp.Message
	err = ReceivePfcpMessage(&pfcpMessage, srcAddr, dstAddr)
	assert.Nil(t, err)

	assert.Equal(t, testPfcpRsp, pfcpMessage)
}
