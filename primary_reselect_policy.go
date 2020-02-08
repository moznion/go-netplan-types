package netplan

type PrimaryReselectPolicy struct {
	val        primaryReselectPolicy
	isAssigned bool
}

type primaryReselectPolicy string

const (
	alwaysPrimaryReselectPolicy  primaryReselectPolicy = "always"
	betterPrimaryReselectPolicy                        = "better"
	failurePrimaryReselectPolicy                       = "failure"
)

func AlwaysPrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        alwaysPrimaryReselectPolicy,
		isAssigned: true,
	}
}
func BetterPrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        betterPrimaryReselectPolicy,
		isAssigned: true,
	}
}

func FailurePrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        failurePrimaryReselectPolicy,
		isAssigned: true,
	}
}

func (prp *PrimaryReselectPolicy) MarshalYAML() (interface{}, error) {
	if prp.isAssigned {
		return prp.val, nil
	}
	return nil, nil
}

func (prp *PrimaryReselectPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val primaryReselectPolicy
	if err := unmarshal(&val); err != nil {
		return err
	}

	prp.val = val
	prp.isAssigned = true

	return nil
}
