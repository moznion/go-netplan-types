package netplan

// NillableUint8 is a data type for nillable uint8 value.
type NillableUint8 struct {
	val        uint8
	isAssigned bool
}

// NillableUint8Of makes a non-nil value with given uint8.
func NillableUint8Of(val uint8) *NillableUint8 {
	return &NillableUint8{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableUint8 as YAML.
// This method used on marshaling YAML internally.
func (nu32 *NillableUint8) MarshalYAML() (interface{}, error) {
	if nu32.isAssigned {
		return nu32.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableUint8 as YAML.
// This method used on unmarshaling YAML internally.
func (nu32 *NillableUint8) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint8
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu32.val = val
	nu32.isAssigned = true

	return nil
}
