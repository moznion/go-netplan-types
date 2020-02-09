package netplan

type NilableString struct {
	val        string
	isAssigned bool
}

func NilableStringOf(val string) *NilableString {
	return &NilableString{
		val:        val,
		isAssigned: true,
	}
}

func (ns *NilableString) MarshalYAML() (interface{}, error) {
	if ns.isAssigned {
		return ns.val, nil
	}
	return nil, nil
}

func (ns *NilableString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}

	ns.val = val
	ns.isAssigned = true

	return nil
}
