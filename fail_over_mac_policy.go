package netplan

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

func NoneFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        noneFailOverMACPolicy,
		isAssigned: true,
	}
}

func ActiveFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        activeFailOverMACPolicy,
		isAssigned: true,
	}
}

func FollowFailOverMACPolicy() *FailOverMACPolicy {
	return &FailOverMACPolicy{
		val:        followFailOverMACPolicy,
		isAssigned: true,
	}
}

func (fomp *FailOverMACPolicy) MarshalYAML() (interface{}, error) {
	if fomp.isAssigned {
		return fomp.val, nil
	}
	return nil, nil
}

func (fomp *FailOverMACPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val failOverMACPolicy
	if err := unmarshal(&val); err != nil {
		return err
	}

	fomp.val = val
	fomp.isAssigned = true

	return nil
}
