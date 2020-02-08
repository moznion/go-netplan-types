package netplan

type ARPValidate struct {
	val        arpValidate
	isAssigned bool
}

type arpValidate string

const (
	noneArpValidate   arpValidate = "none"
	activeArpValidate             = "active"
	backupArpValidate             = "backup"
	allArpValidate                = "all"
)

func NoneArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        noneArpValidate,
		isAssigned: true,
	}
}

func ActiveArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        activeArpValidate,
		isAssigned: true,
	}
}

func BackupArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        backupArpValidate,
		isAssigned: true,
	}
}

func AllArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        allArpValidate,
		isAssigned: true,
	}
}

func (av *ARPValidate) MarshalYAML() (interface{}, error) {
	if av.isAssigned {
		return av.val, nil
	}
	return nil, nil
}

func (av *ARPValidate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val arpValidate
	if err := unmarshal(&val); err != nil {
		return err
	}

	av.val = val
	av.isAssigned = true

	return nil
}
