package netplan

type ARPAllTargets struct {
	val        arpAllTargets
	isAssigned bool
}

type arpAllTargets string

const (
	anyARPAllTargets arpAllTargets = "any"
	allARPAllTargets               = "all"
)

func AnyARPAllTargets() *ARPAllTargets {
	return &ARPAllTargets{
		val:        anyARPAllTargets,
		isAssigned: true,
	}
}

func AllARPAllTargets() *ARPAllTargets {
	return &ARPAllTargets{
		val:        allARPAllTargets,
		isAssigned: true,
	}
}

func (aat *ARPAllTargets) MarshalYAML() (interface{}, error) {
	if aat.isAssigned {
		return aat.val, nil
	}
	return nil, nil
}

func (aat *ARPAllTargets) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val arpAllTargets
	if err := unmarshal(&val); err != nil {
		return err
	}

	aat.val = val
	aat.isAssigned = true

	return nil
}
