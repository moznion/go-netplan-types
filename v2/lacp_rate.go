package netplan

// LACPRate represents ntplan's lacp-rate as nillable.
type LACPRate struct {
	val        lacpRate
	isAssigned bool
}

type lacpRate string

const (
	slowLACPRate lacpRate = "slow"
	fastLACPRate          = "fast"
)

// SlowLACPRate returns `slow` lacp-rate.
func SlowLACPRate() *LACPRate {
	return &LACPRate{
		val:        slowLACPRate,
		isAssigned: true,
	}
}

// FastLACPRate returns `fast` lacp-rate.
func FastLACPRate() *LACPRate {
	return &LACPRate{
		val:        fastLACPRate,
		isAssigned: true,
	}
}

// MarshalYAML marshals LACPRate as YAML.
// This method used on marshaling YAML internally.
func (lr *LACPRate) MarshalYAML() (interface{}, error) {
	if lr.isAssigned {
		return lr.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals LACPRate as YAML.
// This method used on unmarshaling YAML internally.
func (lr *LACPRate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val lacpRate
	if err := unmarshal(&val); err != nil {
		return err
	}

	lr.val = val
	lr.isAssigned = true

	return nil
}
