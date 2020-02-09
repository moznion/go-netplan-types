package netplan

// FailOverMACPolicy represents netplan's fail-over-mac-policy as nillable.
type FailOverMACPolicy struct {
	val        failOverMACPolicy
	isAssigned bool
}

type failOverMACPolicy string

const (
	noneFailOverMACPolicy   failOverMACPolicy = "none"
	activeFailOverMACPolicy                   = "active"
	followFailOverMACPolicy                   = "follow"
)

// NoneFailOverMACPolicy returns `none` fail-over-mac-policy.
func NoneFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        noneFailOverMACPolicy,
		isAssigned: true,
	}
}

// ActiveFailOverMACPolicy returns `active` fail-over-mac-policy.
func ActiveFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        activeFailOverMACPolicy,
		isAssigned: true,
	}
}

// FollowFailOverMACPolicy returns `follow` fail-over-mac-policy.
func FollowFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        followFailOverMACPolicy,
		isAssigned: true,
	}
}

// MarshalYAML marshals FailOverMACPolicy as YAML.
// This method used on marshaling YAML internally.
func (fomp *FailOverMACPolicy) MarshalYAML() (interface{}, error) {
	if fomp.isAssigned {
		return fomp.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals FailOverMACPolicy as YAML.
// This method used on unmarshaling YAML internally.
func (fomp *FailOverMACPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val failOverMACPolicy
	if err := unmarshal(&val); err != nil {
		return err
	}

	fomp.val = val
	fomp.isAssigned = true

	return nil
}
