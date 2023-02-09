package pfcpType

type Multiplier struct {
	Multiplierdata []byte
}

func (m *Multiplier) MarshalBinary() (data []byte, err error) {
	data = m.Multiplierdata

	return data, nil
}

func (m *Multiplier) UnmarshalBinary(data []byte) error {
	m.Multiplierdata = data

	return nil
}
