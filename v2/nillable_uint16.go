package netplan

// NillableUint16 is a data type for nillable uint16 value.
type NillableUint16 struct {
	val        uint16
	isAssigned bool
}

// NillableUint16Of makes a non-nil value with given uint16.
func NillableUint16Of(val uint16) *NillableUint16 {
	return &NillableUint16{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableUint16 as YAML.
// This method used on marshaling YAML internally.
func (nu16 *NillableUint16) MarshalYAML() (interface{}, error) {
	if nu16.isAssigned {
		return nu16.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableUint16 as YAML.
// This method used on unmarshaling YAML internally.
func (nu16 *NillableUint16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint16
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu16.val = val
	nu16.isAssigned = true

	return nil
}
