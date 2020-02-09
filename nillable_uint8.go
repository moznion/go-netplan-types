package go_netplan_types

type NillableUint8 struct {
	val        uint8
	isAssigned bool
}

func NillableUint8Of(val uint8) *NillableUint8 {
	return &NillableUint8{
		val:        val,
		isAssigned: true,
	}
}

func (nu32 *NillableUint8) MarshalYAML() (interface{}, error) {
	if nu32.isAssigned {
		return nu32.val, nil
	}
	return nil, nil
}

func (nu32 *NillableUint8) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint8
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu32.val = val
	nu32.isAssigned = true

	return nil
}
