package pfcpType

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalNodeID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nodeID   NodeID
		expected []byte
	}{
		{
			NodeID{
				NodeIdType: NodeIdTypeIpv4Address,
				IP:         net.ParseIP("192.168.0.1").To4(),
			},
			[]byte{0x0, 0xC0, 0xA8, 0x0, 0x1},
		},
		{
			NodeID{
				NodeIdType: NodeIdTypeIpv6Address,
				IP:         net.ParseIP("2001:db8:0000:1:1:1:1:1"),
			},
			[]byte{0x1, 0x20, 0x1, 0xD, 0xB8, 0x0, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1},
		},
		{
			NodeID{
				NodeIdType: NodeIdTypeFqdn,
				FQDN:       "www.example.com",
			},
			[]byte{
				0x2,
				0x3, 0x77, 0x77, 0x77,
				0x7, 0x65, 0x78, 0x61, 0x6D, 0x70, 0x6C, 0x65,
				0x3, 0x63, 0x6F, 0x6D,
			},
		},
		{
			NodeID{
				NodeIdType: NodeIdTypeFqdn,
				FQDN:       "UPF-a.free5gc.org",
			}, // case-insnesitive case
			[]byte{
				0x2,
				0x5, 0x55, 0x50, 0x46, 0x2d, 0x61,
				0x7, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x3, 0x6F, 0x72, 0x67,
			},
		},
	}

	for _, test := range tests {
		buf, err := test.nodeID.MarshalBinary()
		require.NoError(t, err)
		require.Equal(t, test.expected, buf)
	}
}

func TestUnmarshalNodeID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputbuf []byte
		expected NodeID
	}{
		{
			[]byte{0x0, 0xC0, 0xA8, 0x0, 0x1},
			NodeID{
				NodeIdType: NodeIdTypeIpv4Address,
				IP:         net.ParseIP("192.168.0.1").To4(),
			},
		},
		{
			[]byte{0x1, 0x20, 0x1, 0xD, 0xB8, 0x0, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x1},
			NodeID{
				NodeIdType: NodeIdTypeIpv6Address,
				IP:         net.ParseIP("2001:db8:0000:1:1:1:1:1"),
			},
		},
		{
			[]byte{
				0x2,
				0x3, 0x77, 0x77, 0x77,
				0x7, 0x65, 0x78, 0x61, 0x6D, 0x70, 0x6C, 0x65,
				0x3, 0x63, 0x6F, 0x6D,
			},
			NodeID{
				NodeIdType: NodeIdTypeFqdn,
				FQDN:       "www.example.com",
			},
		},
		{
			[]byte{
				0x2,
				0x5, 0x55, 0x50, 0x46, 0x2d, 0x61,
				0x7, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x3, 0x6F, 0x72, 0x67,
			},
			NodeID{
				NodeIdType: NodeIdTypeFqdn,
				FQDN:       "UPF-a.free5gc.org",
			},
		},
	}

	var nodeID NodeID
	for _, test := range tests {
		err := nodeID.UnmarshalBinary(test.inputbuf)
		require.NoError(t, err)
		require.Equal(t, test.expected, nodeID)
	}
}
