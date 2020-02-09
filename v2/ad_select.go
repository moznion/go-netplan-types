package netplan

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

func StableAdSelect() *AdSelect {
	return &AdSelect{
		val:        stableAdSelect,
		isAssigned: true,
	}
}

func BandwidthAdSelect() *AdSelect {
	return &AdSelect{
		val:        bandwidthAdSelect,
		isAssigned: true,
	}
}

func CountAdSelect() *AdSelect {
	return &AdSelect{
		val:        countAdSelect,
		isAssigned: true,
	}
}

func (as *AdSelect) MarshalYAML() (interface{}, error) {
	if as.isAssigned {
		return as.val, nil
	}
	return nil, nil
}

func (as *AdSelect) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val adSelect
	if err := unmarshal(&val); err != nil {
		return err
	}

	as.val = val
	as.isAssigned = true

	return nil
}
