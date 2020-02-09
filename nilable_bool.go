package go_netplan_types

type NilableBool struct {
	val        bool
	isAssigned bool
}

func NilableBoolOf(val bool) *NilableBool {
	return &NilableBool{
		val:        val,
		isAssigned: true,
	}
}

func (nb *NilableBool) MarshalYAML() (interface{}, error) {
	if nb.isAssigned {
		return nb.val, nil
	}
	return nil, nil
}

func (nb *NilableBool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bool
	if err := unmarshal(&val); err != nil {
		return err
	}

	nb.val = val
	nb.isAssigned = true

	return nil
}
