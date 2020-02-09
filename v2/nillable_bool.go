package netplan

// NillableBool is a data type for nillable bool value.
type NillableBool struct {
	val        bool
	isAssigned bool
}

// NillableBoolOf makes a non-nil value with given bool.
func NillableBoolOf(val bool) *NillableBool {
	return &NillableBool{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableBool as YAML.
// This method used on marshaling YAML internally.
func (nb *NillableBool) MarshalYAML() (interface{}, error) {
	if nb.isAssigned {
		return nb.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableBool as YAML.
// This method used on unmarshaling YAML internally.
func (nb *NillableBool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bool
	if err := unmarshal(&val); err != nil {
		return err
	}

	nb.val = val
	nb.isAssigned = true

	return nil
}
