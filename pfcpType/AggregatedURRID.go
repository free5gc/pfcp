package pfcpType

import (
	"encoding/binary"
)

type AggregatedURRID struct {
	AggregatedURRIDdata uint32
}

func (a *AggregatedURRID) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, a.AggregatedURRIDdata)

	return data, nil
}

func (a *AggregatedURRID) UnmarshalBinary(data []byte) error {
	a.AggregatedURRIDdata = binary.BigEndian.Uint32(data)

	return nil
}
