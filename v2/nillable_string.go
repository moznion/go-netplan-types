package netplan

// NillableString is a data type for nillable string value.
type NillableString struct {
	val        string
	isAssigned bool
}

// NillableStringOf makes a non-nil value with given string.
func NillableStringOf(val string) *NillableString {
	return &NillableString{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableString as YAML.
// This method used on marshaling YAML internally.
func (ns *NillableString) MarshalYAML() (interface{}, error) {
	if ns.isAssigned {
		return ns.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableString as YAML.
// This method used on unmarshaling YAML internally.
func (ns *NillableString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}

	ns.val = val
	ns.isAssigned = true

	return nil
}
