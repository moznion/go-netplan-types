package go_netplan_types

type NilableUint32 struct {
	val        uint32
	isAssigned bool
}

func NilableUint32Of(val uint32) *NilableUint32 {
	return &NilableUint32{
		val:        val,
		isAssigned: true,
	}
}

func (nu32 *NilableUint32) MarshalYAML() (interface{}, error) {
	if nu32.isAssigned {
		return nu32.val, nil
	}
	return nil, nil
}

func (nu32 *NilableUint32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint32
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu32.val = val
	nu32.isAssigned = true

	return nil
}
