package go_netplan_types

type NilableUint16 struct {
	val        uint16
	isAssigned bool
}

func NilableUint16Of(val uint16) *NilableUint16 {
	return &NilableUint16{
		val:        val,
		isAssigned: true,
	}
}

func (nu16 *NilableUint16) MarshalYAML() (interface{}, error) {
	if nu16.isAssigned {
		return nu16.val, nil
	}
	return nil, nil
}

func (nu16 *NilableUint16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint16
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu16.val = val
	nu16.isAssigned = true

	return nil
}
