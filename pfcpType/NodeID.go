package pfcpType

import (
	"errors"
	"fmt"
	"net"

	"github.com/free5gc/pfcp/logger"
)

const (
	NodeIdTypeIpv4Address uint8 = iota
	NodeIdTypeIpv6Address
	NodeIdTypeFqdn
)

type NodeID struct {
	NodeIdType uint8 // 0x00001111
	IP         net.IP
	FQDN       string
}

func (n *NodeID) MarshalBinary() ([]byte, error) {
	var data []byte
	data = append([]byte(""), n.NodeIdType)

	var nodeIdValue []byte

	// Octet 6 to o
	switch n.NodeIdType {
	case NodeIdTypeIpv4Address:
		if len(n.IP) == 0 || n.FQDN != "" {
			return []byte(""), errors.New("type of node ID is ipv4, should fill IP field and FQDN field should be empty")
		}
		if len(n.IP) != net.IPv4len {
			return []byte(""), errors.New("length of node id data shall be 4 Octet if node id is an IPv4 address")
		}
		nodeIdValue = n.IP

	case NodeIdTypeIpv6Address:
		if len(n.IP) == 0 || n.FQDN != "" {
			return []byte(""), errors.New("type of node ID is ipv6, should fill IP field and FQDN field should be empty")
		}
		if len(n.IP) != net.IPv6len {
			return []byte(""), errors.New("length of node id data shall be 16 Octet if node id is an IPv6 address")
		}
		nodeIdValue = n.IP

	case NodeIdTypeFqdn:
		if n.FQDN == "" || n.IP != nil {
			return []byte(""), errors.New("type of node ID is fqdn, should fill FQDN field and IP field should be empty")
		}
		if rfc1035RR, err := fqdnToRfc1035(n.FQDN, false); err != nil {
			return nil, err
		} else {
			nodeIdValue = rfc1035RR
		}
	default:
		return nil, errors.New("type of node id should be ipv4, ipv6 or fqdn")
	}

	data = append(data, nodeIdValue...)

	return data, nil
}

func (n *NodeID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	n.NodeIdType = data[idx] & Mask4
	idx = idx + 1

	// Octet 6 to o
	switch n.NodeIdType {
	case NodeIdTypeIpv4Address:
		if length < idx+net.IPv4len {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		n.IP = data[idx : idx+net.IPv4len]
		n.FQDN = ""
		idx = idx + net.IPv4len
	case NodeIdTypeIpv6Address:
		if length < idx+net.IPv6len {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		n.IP = data[idx : idx+net.IPv6len]
		n.FQDN = ""
		idx = idx + net.IPv6len
	case NodeIdTypeFqdn:
		rfc1035RR := data[idx:]
		n.FQDN = rfc1035tofqdn(rfc1035RR)
		n.IP = nil
		idx = idx + uint16(len(rfc1035RR))
	}

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	return nil
}

func (n *NodeID) ResolveNodeIdToIp() net.IP {
	switch n.NodeIdType {
	case NodeIdTypeIpv4Address, NodeIdTypeIpv6Address:
		return n.IP
	case NodeIdTypeFqdn:
		if ns, err := net.LookupHost(n.FQDN); err != nil {
			logger.PFCPLog.Warnf("Host lookup failed: %+v", err)
			return net.IPv4zero
		} else {
			return net.ParseIP(ns[0])
		}
	default:
		return net.IPv4zero
	}
}
