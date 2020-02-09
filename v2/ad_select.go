package netplan

// AdSelect represents netplan's `ad-select` as nillable.
type AdSelect struct {
	val        adSelect
	isAssigned bool
}

type adSelect string

const (
	stableAdSelect    adSelect = "stable"
	bandwidthAdSelect          = "bandwidth"
	countAdSelect              = "count"
)

// StableAdSelect returns `stable` ad-select.
func StableAdSelect() *AdSelect {
	return &AdSelect{
		val:        stableAdSelect,
		isAssigned: true,
	}
}

// BandwidthAdSelect returns `bandwidth` ad-select.
func BandwidthAdSelect() *AdSelect {
	return &AdSelect{
		val:        bandwidthAdSelect,
		isAssigned: true,
	}
}

// CountAdSelect returns `count` ad-select.
func CountAdSelect() *AdSelect {
	return &AdSelect{
		val:        countAdSelect,
		isAssigned: true,
	}
}

// MarshalYAML marshals AdSelect as YAML.
// This method used on marshaling YAML internally.
func (as *AdSelect) MarshalYAML() (interface{}, error) {
	if as.isAssigned {
		return as.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals AdSelect as YAML.
// This method used on unmarshaling YAML internally.
func (as *AdSelect) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val adSelect
	if err := unmarshal(&val); err != nil {
		return err
	}

	as.val = val
	as.isAssigned = true

	return nil
}
