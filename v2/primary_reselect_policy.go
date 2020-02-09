package netplan

// PrimaryReselectPolicy represents netplan's primary-reselect-policy as nillable.
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

// AlwaysPrimaryReselectPolicy returns `always` primary-reselect-policy.
func AlwaysPrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        alwaysPrimaryReselectPolicy,
		isAssigned: true,
	}
}

// BetterPrimaryReselectPolicy returns `better` primary-reselect-policy.
func BetterPrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        betterPrimaryReselectPolicy,
		isAssigned: true,
	}
}

// FailurePrimaryReselectPolicy returns `failure` primary-reselect-policy.
func FailurePrimaryReselectPolicy() *PrimaryReselectPolicy {
	return &PrimaryReselectPolicy{
		val:        failurePrimaryReselectPolicy,
		isAssigned: true,
	}
}

// MarshalYAML marshals PrimaryReselectPolicy as YAML.
// This method used on marshaling YAML internally.
func (prp *PrimaryReselectPolicy) MarshalYAML() (interface{}, error) {
	if prp.isAssigned {
		return prp.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals PrimaryReselectPolicy as YAML.
// This method used on unmarshaling YAML internally.
func (prp *PrimaryReselectPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val primaryReselectPolicy
	if err := unmarshal(&val); err != nil {
		return err
	}

	prp.val = val
	prp.isAssigned = true

	return nil
}
