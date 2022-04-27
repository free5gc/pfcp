package pfcp

import "net"

type ReceiveEvent struct {
	Type       ReceiveEventType
	RemoteAddr *net.UDPAddr
	RcvMsg     *Message
}

type ReceiveEventType uint8

const (
	ReceiveEventTypeResendRequest ReceiveEventType = iota
	ReceiveEventTypeValidResponse
)
