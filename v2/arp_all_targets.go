package netplan

// ARPAllTargets represents netplan's `arp-all-targets` as nillable.
type ARPAllTargets struct {
	val        arpAllTargets
	isAssigned bool
}

type arpAllTargets string

const (
	anyARPAllTargets arpAllTargets = "any"
	allARPAllTargets               = "all"
)

// AnyARPAllTargets returns `any` arp-all-targets.
func AnyARPAllTargets() *ARPAllTargets {
	return &ARPAllTargets{
		val:        anyARPAllTargets,
		isAssigned: true,
	}
}

// AllARPAllTargets returns `all` arp-all-targets.
func AllARPAllTargets() *ARPAllTargets {
	return &ARPAllTargets{
		val:        allARPAllTargets,
		isAssigned: true,
	}
}

// MarshalYAML marshals ARPAllTargets  as YAML.
// This method used on marshaling YAML internally.
func (aat *ARPAllTargets) MarshalYAML() (interface{}, error) {
	if aat.isAssigned {
		return aat.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals ARPAllTargets as YAML.
// This method used on unmarshaling YAML internally.
func (aat *ARPAllTargets) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val arpAllTargets
	if err := unmarshal(&val); err != nil {
		return err
	}

	aat.val = val
	aat.isAssigned = true

	return nil
}
