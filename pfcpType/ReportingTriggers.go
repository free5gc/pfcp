package pfcpType

import (
	"fmt"
)

type ReportingTriggers struct {
	Liusa bool
	Droth bool
	Stopt bool
	Start bool
	Quhti bool
	Timth bool
	Volth bool
	Perio bool
	Evequ bool
	Eveth bool
	Macar bool
	Envcl bool
	Timqu bool
	Volqu bool
	Ipmjl bool
	Quvti bool
}

func (r *ReportingTriggers) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(r.Liusa)<<7 |
		btou(r.Droth)<<6 |
		btou(r.Stopt)<<5 |
		btou(r.Start)<<4 |
		btou(r.Quhti)<<3 |
		btou(r.Timth)<<2 |
		btou(r.Volth)<<1 |
		btou(r.Perio)
	data = append([]byte(""), tmpUint8)

	// Octet 6
	tmpUint8 = btou(r.Quvti)<<7 |
		btou(r.Ipmjl)<<6 |
		btou(r.Evequ)<<5 |
		btou(r.Eveth)<<4 |
		btou(r.Macar)<<3 |
		btou(r.Envcl)<<2 |
		btou(r.Timqu)<<1 |
		btou(r.Volqu)
	data = append(data, tmpUint8)

	return data, nil
}

func (r *ReportingTriggers) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	r.Liusa = utob(data[idx] & BitMask8)
	r.Droth = utob(data[idx] & BitMask7)
	r.Stopt = utob(data[idx] & BitMask6)
	r.Start = utob(data[idx] & BitMask5)
	r.Quhti = utob(data[idx] & BitMask4)
	r.Timth = utob(data[idx] & BitMask3)
	r.Volth = utob(data[idx] & BitMask2)
	r.Perio = utob(data[idx] & BitMask1)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	r.Quvti = utob(data[idx] & BitMask8)
	r.Ipmjl = utob(data[idx] & BitMask7)
	r.Evequ = utob(data[idx] & BitMask6)
	r.Eveth = utob(data[idx] & BitMask5)
	r.Macar = utob(data[idx] & BitMask4)
	r.Envcl = utob(data[idx] & BitMask3)
	r.Timqu = utob(data[idx] & BitMask2)
	r.Volqu = utob(data[idx] & BitMask1)
	idx = idx + 1

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
