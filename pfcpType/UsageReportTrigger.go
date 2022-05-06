package pfcpType

import "fmt"

type UsageReportTrigger struct {
	Immer bool
	Droth bool
	Stopt bool
	Start bool
	Quhti bool
	Timth bool
	Volth bool
	Perio bool
	Eveth bool
	Macar bool
	Envcl bool
	Monit bool
	Termr bool
	Liusa bool
	Timqu bool
	Volqu bool
	Emrre bool
	Quvti bool
	Ipmjl bool
	Tebur bool
	Evequ bool
}

func (u *UsageReportTrigger) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(u.Immer)<<7 |
		btou(u.Droth)<<6 |
		btou(u.Stopt)<<5 |
		btou(u.Start)<<4 |
		btou(u.Quhti)<<3 |
		btou(u.Timth)<<2 |
		btou(u.Volth)<<1 |
		btou(u.Perio)
	data = append(data, tmpUint8)

	// Octet 6
	tmpUint8 = btou(u.Eveth)<<7 |
		btou(u.Macar)<<6 |
		btou(u.Envcl)<<5 |
		btou(u.Monit)<<4 |
		btou(u.Termr)<<3 |
		btou(u.Liusa)<<2 |
		btou(u.Timqu)<<1 |
		btou(u.Volqu)
	data = append(data, tmpUint8)

	// Octet 7
	tmpUint8 = btou(u.Emrre)<<4 |
		btou(u.Quvti)<<3 |
		btou(u.Ipmjl)<<2 |
		btou(u.Tebur)<<1 |
		btou(u.Evequ)
	data = append(data, tmpUint8)

	return data, nil
}

func (u *UsageReportTrigger) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length != 3 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	u.Immer = utob(data[idx] & BitMask8)
	u.Droth = utob(data[idx] & BitMask7)
	u.Stopt = utob(data[idx] & BitMask6)
	u.Start = utob(data[idx] & BitMask5)
	u.Quhti = utob(data[idx] & BitMask4)
	u.Timth = utob(data[idx] & BitMask3)
	u.Volth = utob(data[idx] & BitMask2)
	u.Perio = utob(data[idx] & BitMask1)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	u.Eveth = utob(data[idx] & BitMask8)
	u.Macar = utob(data[idx] & BitMask7)
	u.Envcl = utob(data[idx] & BitMask6)
	u.Monit = utob(data[idx] & BitMask5)
	u.Termr = utob(data[idx] & BitMask4)
	u.Liusa = utob(data[idx] & BitMask3)
	u.Timqu = utob(data[idx] & BitMask2)
	u.Volqu = utob(data[idx] & BitMask1)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	u.Emrre = utob(data[idx] & BitMask5)
	u.Quvti = utob(data[idx] & BitMask4)
	u.Ipmjl = utob(data[idx] & BitMask3)
	u.Tebur = utob(data[idx] & BitMask2)
	u.Evequ = utob(data[idx] & BitMask1)
	idx = idx + 1

	if length != idx {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
