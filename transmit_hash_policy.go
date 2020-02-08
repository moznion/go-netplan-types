package netplan

type TransmitHashPolicy struct {
	val        transmitHashPolicy
	isAssigned bool
}

type transmitHashPolicy string

const (
	layer2TransmitHashPolicy     transmitHashPolicy = "layer2"
	layer3And4TransmitHashPolicy                    = "layer3+4"
	layer2And3TransmitHashPolicy                    = "layer2+3"
	encap2And3TransmitHashPolicy                    = "encap2+3"
	encap3And4TransmitHashPolicy                    = "encap3+4"
)

func Layer2TransmitHashPolicy() *TransmitHashPolicy {
	return &TransmitHashPolicy{
		val:        layer2TransmitHashPolicy,
		isAssigned: true,
	}
}

func Layer3And4TransmitHashPolicy() *TransmitHashPolicy {
	return &TransmitHashPolicy{
		val:        layer3And4TransmitHashPolicy,
		isAssigned: true,
	}
}

func Layer2And3TransmitHashPolicy() *TransmitHashPolicy {
	return &TransmitHashPolicy{
		val:        layer2And3TransmitHashPolicy,
		isAssigned: true,
	}
}

func Encap2And3TransmitHashPolicy() *TransmitHashPolicy {
	return &TransmitHashPolicy{
		val:        encap2And3TransmitHashPolicy,
		isAssigned: true,
	}
}

func Encap3And4TransmitHashPolicy() *TransmitHashPolicy {
	return &TransmitHashPolicy{
		val:        encap3And4TransmitHashPolicy,
		isAssigned: true,
	}
}

func (thp *TransmitHashPolicy) MarshalYAML() (interface{}, error) {
	if thp.isAssigned {
		return thp.val, nil
	}
	return nil, nil
}

func (thp *TransmitHashPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val transmitHashPolicy
	if err := unmarshal(&val); err != nil {
		return err
	}

	thp.val = val
	thp.isAssigned = true

	return nil
}
