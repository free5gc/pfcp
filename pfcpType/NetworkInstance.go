package pfcpType

type NetworkInstance struct {
	NetworkInstance string
	FQDNEncoding    bool
}

func (n *NetworkInstance) MarshalBinary() ([]byte, error) {
	if n.FQDNEncoding {
		return fqdnToRfc1035(n.NetworkInstance, true)
	}
	return []byte(n.NetworkInstance), nil
}

func (n *NetworkInstance) UnmarshalBinary(data []byte) error {
	if n.FQDNEncoding {
		n.NetworkInstance = rfc1035tofqdn(data)
	} else {
		n.NetworkInstance = string(data[:])
	}
	return nil
}
