package go_netplan_types

type NillableString struct {
	val        string
	isAssigned bool
}

func NillableStringOf(val string) *NillableString {
	return &NillableString{
		val:        val,
		isAssigned: true,
	}
}

func (ns *NillableString) MarshalYAML() (interface{}, error) {
	if ns.isAssigned {
		return ns.val, nil
	}
	return nil, nil
}

func (ns *NillableString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}

	ns.val = val
	ns.isAssigned = true

	return nil
}
