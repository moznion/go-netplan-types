package netplan

type NillableBool struct {
	val        bool
	isAssigned bool
}

func NillableBoolOf(val bool) *NillableBool {
	return &NillableBool{
		val:        val,
		isAssigned: true,
	}
}

func (nb *NillableBool) MarshalYAML() (interface{}, error) {
	if nb.isAssigned {
		return nb.val, nil
	}
	return nil, nil
}

func (nb *NillableBool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bool
	if err := unmarshal(&val); err != nil {
		return err
	}

	nb.val = val
	nb.isAssigned = true

	return nil
}
