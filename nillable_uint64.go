package go_netplan_types

type NillableUint64 struct {
	val        uint64
	isAssigned bool
}

func NillableUint64Of(val uint64) *NillableUint64 {
	return &NillableUint64{
		val:        val,
		isAssigned: true,
	}
}

func (nu64 *NillableUint64) MarshalYAML() (interface{}, error) {
	if nu64.isAssigned {
		return nu64.val, nil
	}
	return nil, nil
}

func (nu64 *NillableUint64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint64
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu64.val = val
	nu64.isAssigned = true

	return nil
}
