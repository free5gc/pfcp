package pfcpType

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/bits"
)

type GBR struct {
	ULGBR uint64 // 40-bit data
	DLGBR uint64 // 40-bit data
}

func (g *GBR) MarshalBinary() ([]byte, error) {
	buf := &bytes.Buffer{}

	if bits.Len64(g.ULGBR) > 40 {
		return nil, fmt.Errorf("UL GBR shall not be greater than 40 bits binary integer")
	}
	if bits.Len64(g.DLGBR) > 40 {
		return nil, fmt.Errorf("DL GBR shall not be greater than 40 bits binary integer")
	}

	gbrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(gbrBytes, g.ULGBR)

	if err := binary.Write(buf, binary.BigEndian, gbrBytes[3:]); err != nil {
		return nil, fmt.Errorf("write ULGBR fail: %s", err)
	}

	binary.BigEndian.PutUint64(gbrBytes, g.DLGBR)

	if err := binary.Write(buf, binary.BigEndian, gbrBytes[3:]); err != nil {
		return nil, fmt.Errorf("write DLGBR fail: %s", err)
	}

	return buf.Bytes(), nil
}

func (g *GBR) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)

	GBRBytes := make([]byte, 5)
	uint64Byte := make([]byte, 8)

	if err := binary.Read(buf, binary.BigEndian, GBRBytes); err != nil {
		return fmt.Errorf("read UL GBR fail: %s", err)
	}

	copy(uint64Byte[3:], GBRBytes)
	g.ULGBR = binary.BigEndian.Uint64(uint64Byte)

	if err := binary.Read(buf, binary.BigEndian, GBRBytes); err != nil {
		return fmt.Errorf("read DL GBR fail: %s", err)
	}

	copy(uint64Byte[3:], GBRBytes)
	g.DLGBR = binary.BigEndian.Uint64(uint64Byte)

	return nil
}
