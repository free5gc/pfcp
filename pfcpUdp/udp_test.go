package pfcpUdp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFromContinuesAfterMalformedDatagram(t *testing.T) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	require.NoError(t, err)
	defer conn.Close()

	server := &PfcpServer{Conn: conn}
	client, err := net.DialUDP("udp", nil, conn.LocalAddr().(*net.UDPAddr))
	require.NoError(t, err)
	defer client.Close()

	_, err = client.Write([]byte{0x20, 0x05, 0x00, 0x00})
	require.NoError(t, err)
	_, err = server.ReadFrom()
	require.Error(t, err)

	validHeartbeatRequest := []byte{0x20, 0x01, 0x00, 0x04, 0x00, 0x00, 0x01, 0x00}
	_, err = client.Write(validHeartbeatRequest)
	require.NoError(t, err)

	msg, err := server.ReadFrom()
	require.NoError(t, err)
	require.NotNil(t, msg)
}
