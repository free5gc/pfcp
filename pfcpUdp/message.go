package pfcpUdp

import (
	"net"

	"github.com/free5gc/pfcp"
)

type Message struct {
	RemoteAddr  *net.UDPAddr
	PfcpMessage *pfcp.Message
}

func NewMessage(remoteAddr *net.UDPAddr, pfcpMessage *pfcp.Message) (msg *Message) {
	return &Message{
		RemoteAddr:  remoteAddr,
		PfcpMessage: pfcpMessage,
	}
}

func (m *Message) MessageType() pfcp.MessageType {
	return m.PfcpMessage.Header.MessageType
}
