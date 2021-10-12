package pfcpType

type NetworkInstance struct {
	NetworkInstance string
}

func (n *NetworkInstance) MarshalBinary() ([]byte, error) {
	return fqdnToRfc1035(n.NetworkInstance)
}

func (n *NetworkInstance) UnmarshalBinary(data []byte) error {
	n.NetworkInstance = rfc1035tofqdn(data)
	return nil
}
