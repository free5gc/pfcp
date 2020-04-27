package pfcpUdp

import (
	"net"

	"free5gc/lib/pfcp"
)

const (
	PFCP_PORT        = 8805
	PFCP_MAX_UDP_LEN = 2048
)

type PfcpServer struct {
	Addr string
	Conn *net.UDPConn
}

func NewPfcpServer(addr string) (PfcpServer, error) {
	var server PfcpServer
	server.Addr = addr
	err := server.Listen()

	return server, err
}

func (p *PfcpServer) Listen() error {
	var serverIp net.IP
	if p.Addr == "" {
		serverIp = net.IPv4zero
	} else {
		serverIp = net.ParseIP(p.Addr)
	}
	addr := &net.UDPAddr{
		IP:   serverIp,
		Port: PFCP_PORT,
	}

	conn, err := net.ListenUDP("udp", addr)
	p.Conn = conn
	return err
}

func (p *PfcpServer) ReadFrom(msg *pfcp.Message) (*net.UDPAddr, error) {
	buf := make([]byte, PFCP_MAX_UDP_LEN)
	n, addr, err := p.Conn.ReadFromUDP(buf)
	if err != nil {
		return addr, err
	}

	err = msg.Unmarshal(buf[:n])
	if err != nil {
		return addr, err
	}

	return addr, nil
}

func (p *PfcpServer) WriteTo(msg pfcp.Message, addr *net.UDPAddr) error {
	buf, err := msg.Marshal()
	if err != nil {
		return err
	}

	/*TODO: check if all bytes of buf are sent*/
	_, err = p.Conn.WriteToUDP(buf, addr)
	if err != nil {
		return err
	}

	return nil
}

func (p *PfcpServer) Close() error {
	return p.Conn.Close()
}

// Send a PFCP message and close UDP connection
func SendPfcpMessage(msg pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf, err := msg.Marshal()
	if err != nil {
		return err
	}

	/*TODO: check if all bytes of buf are sent*/
	_, err = conn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

// Receive a PFCP message and close UDP connection
func ReceivePfcpMessage(msg *pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, PFCP_MAX_UDP_LEN)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

	err = msg.Unmarshal(buf[:n])
	if err != nil {
		return err
	}

	return nil
}
