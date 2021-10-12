package pfcpType

import (
	"fmt"
	"math/bits"
)

type DownlinkDataServiceInformation struct {
	Qfii                        bool
	Ppi                         bool
	PagingPolicyIndicationValue uint8 // 0x00111111
	Qfi                         uint8 // 0x00111111
}

func (d *DownlinkDataServiceInformation) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(d.Qfii)<<1 | btou(d.Ppi)
	data = append([]byte(""), tmpUint8)

	// Octet m
	if d.Ppi {
		if bits.Len8(d.PagingPolicyIndicationValue) > 6 {
			return []byte(""), fmt.Errorf("Paging policy information data shall not be greater than 6 bits binary integer")
		}
		data = append(data, d.PagingPolicyIndicationValue)
	}

	// Octet p
	if d.Qfii {
		if bits.Len8(d.Qfi) > 6 {
			return []byte(""), fmt.Errorf("QFI shall not be greater than 6 bits binary integer")
		}
		data = append(data, d.Qfi)
	}

	return data, nil
}

func (d *DownlinkDataServiceInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	d.Qfii = utob(data[idx] & BitMask2)
	d.Ppi = utob(data[idx] & BitMask1)
	idx = idx + 1

	// Octet m
	if d.Ppi {
		if length < idx+1 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		d.PagingPolicyIndicationValue = data[idx] & Mask6
		idx = idx + 1
	}

	// Octet p
	if d.Qfii {
		if length < idx+1 {
			return fmt.Errorf("Inadequate TLV length: %d", length)
		}
		d.Qfi = data[idx] & Mask6
		idx = idx + 1
	}

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
