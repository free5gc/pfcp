package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalNetworkInstance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in       NetworkInstance
		expected []byte
	}{
		{
			NetworkInstance{
				NetworkInstance: "www.example.com",
			},
			[]byte{
				0x3, 0x77, 0x77, 0x77,
				0x7, 0x65, 0x78, 0x61, 0x6D, 0x70, 0x6C, 0x65,
				0x3, 0x63, 0x6F, 0x6D,
			},
		},
		{
			NetworkInstance{
				NetworkInstance: "UPF-a.free5gc.org",
			}, // case-insnesitive case
			[]byte{
				0x5, 0x55, 0x50, 0x46, 0x2d, 0x61,
				0x7, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x3, 0x6F, 0x72, 0x67,
			},
		},
	}

	for _, test := range tests {
		buf, err := test.in.MarshalBinary()
		require.NoError(t, err)
		require.Equal(t, test.expected, buf)
	}
}

func TestUnmarshalNetworkInstance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputbuf []byte
		expected NetworkInstance
	}{
		{
			[]byte{
				0x3, 0x77, 0x77, 0x77,
				0x7, 0x65, 0x78, 0x61, 0x6D, 0x70, 0x6C, 0x65,
				0x3, 0x63, 0x6F, 0x6D,
			},
			NetworkInstance{
				NetworkInstance: "www.example.com",
			},
		},
		{
			[]byte{
				0x5, 0x55, 0x50, 0x46, 0x2d, 0x61,
				0x7, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x3, 0x6F, 0x72, 0x67,
			},
			NetworkInstance{
				NetworkInstance: "UPF-a.free5gc.org",
			},
		},
	}

	var n NetworkInstance
	for _, test := range tests {
		err := n.UnmarshalBinary(test.inputbuf)
		require.NoError(t, err)
		require.Equal(t, test.expected, n)
	}
}
