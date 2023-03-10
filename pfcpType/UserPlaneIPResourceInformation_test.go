package pfcpType

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUserPlaneIPResourceInformation(t *testing.T) {
	testData := UserPlaneIPResourceInformation{
		Assosi:          true,
		Assoni:          true,
		Teidri:          4,
		V6:              true,
		V4:              true,
		TeidRange:       21,
		Ipv4Address:     net.ParseIP("12.34.56.78").To4(),
		Ipv6Address:     net.ParseIP("2001:db8::68").To16(),
		NetworkInstance: NetworkInstance{NetworkInstance: "free5gc.local"},
		SourceInterface: 12,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{
		0x73, 0x15, 0x0c, 0x22, 0x38, 0x4e, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63, 0x2e, 0x6c, 0x6f,
		0x63, 0x61, 0x6c, 0x0c,
	}, buf)
}

func TestUnmarshalUserPlaneIPResourceInformation(t *testing.T) {
	buf := []byte{
		0x73, 0x15, 0x0c, 0x22, 0x38, 0x4e, 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63, 0x2e, 0x6c, 0x6f,
		0x63, 0x61, 0x6c, 0x0c,
	}
	var testData UserPlaneIPResourceInformation
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UserPlaneIPResourceInformation{
		Assosi:          true,
		Assoni:          true,
		Teidri:          4,
		V6:              true,
		V4:              true,
		TeidRange:       21,
		Ipv4Address:     net.ParseIP("12.34.56.78").To4(),
		Ipv6Address:     net.ParseIP("2001:db8::68").To16(),
		NetworkInstance: NetworkInstance{NetworkInstance: "free5gc.local"},
		SourceInterface: 12,
	}

	assert.Equal(t, expectData, testData)
}
