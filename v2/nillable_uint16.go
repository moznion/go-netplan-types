package netplan

type NillableUint16 struct {
	val        uint16
	isAssigned bool
}

func NillableUint16Of(val uint16) *NillableUint16 {
	return &NillableUint16{
		val:        val,
		isAssigned: true,
	}
}

func (nu16 *NillableUint16) MarshalYAML() (interface{}, error) {
	if nu16.isAssigned {
		return nu16.val, nil
	}
	return nil, nil
}

func (nu16 *NillableUint16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint16
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu16.val = val
	nu16.isAssigned = true

	return nil
}
