package go_netplan_types

type NilableUint64 struct {
	val        uint64
	isAssigned bool
}

func NilableUint64Of(val uint64) *NilableUint64 {
	return &NilableUint64{
		val:        val,
		isAssigned: true,
	}
}

func (nu64 *NilableUint64) MarshalYAML() (interface{}, error) {
	if nu64.isAssigned {
		return nu64.val, nil
	}
	return nil, nil
}

func (nu64 *NilableUint64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint64
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu64.val = val
	nu64.isAssigned = true

	return nil
}
