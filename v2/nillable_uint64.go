package netplan

// NillableUint64 is a data type for nillable uint64 value.
type NillableUint64 struct {
	val        uint64
	isAssigned bool
}

// NillableUint64Of makes a non-nil value with given uint64.
func NillableUint64Of(val uint64) *NillableUint64 {
	return &NillableUint64{
		val:        val,
		isAssigned: true,
	}
}

// MarshalYAML marshals NillableUint64 as YAML.
// This method used on marshaling YAML internally.
func (nu64 *NillableUint64) MarshalYAML() (interface{}, error) {
	if nu64.isAssigned {
		return nu64.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals NillableUint64 as YAML.
// This method used on unmarshaling YAML internally.
func (nu64 *NillableUint64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint64
	if err := unmarshal(&val); err != nil {
		return err
	}

	nu64.val = val
	nu64.isAssigned = true

	return nil
}
