package netplan

// NillableUint32 is a data type for nillable uint32 value.
type NillableUint32 struct {
	val        uint32
	isAssigned bool
}

// NillableUint32Of makes a non-nil value with given uint32.
func NillableUint32Of(val uint32) *NillableUint32 {
	return &NillableUint32{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableUint32 as YAML.
// This method used on marshaling YAML internally.
func (nu32 *NillableUint32) MarshalYAML() (interface{}, error) {
	if nu32.isAssigned {
		return nu32.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableUint32 as YAML.
// This method used on unmarshaling YAML internally.
func (nu32 *NillableUint32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint32
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu32.val = val
	nu32.isAssigned = true

	return nil
}
