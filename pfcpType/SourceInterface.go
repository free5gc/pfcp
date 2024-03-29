package pfcpType

import (
	"fmt"
	"math/bits"
)

const (
	SourceInterfaceAccess uint8 = iota
	SourceInterfaceCore
	SourceInterfaceSgiLanN6Lan
	SourceInterfaceCpFunction
)

type SourceInterface struct {
	InterfaceValue uint8 // 0x00001111
}

func (s *SourceInterface) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if bits.Len8(s.InterfaceValue) > 4 {
		return []byte(""), fmt.Errorf("Interface data shall not be greater than 4 bits binary integer")
	}
	data = append([]byte(""), s.InterfaceValue)

	return data, nil
}

func (s *SourceInterface) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	s.InterfaceValue = data[idx] & Mask4
	idx = idx + 1

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
