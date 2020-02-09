package netplan

// ARPValidate represents netplan's `arp-validate` as nillable.
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

// NoneArpValidate returns `none` arp-validate.
func NoneArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        noneArpValidate,
		isAssigned: true,
	}
}

// ActiveArpValidate returns `active` arp-validate.
func ActiveArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        activeArpValidate,
		isAssigned: true,
	}
}

// BackupArpValidate returns `backup` arp-validate.
func BackupArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        backupArpValidate,
		isAssigned: true,
	}
}

// AllArpValidate returns `all` arp-validate.
func AllArpValidate() *ARPValidate {
	return &ARPValidate{
		val:        allArpValidate,
		isAssigned: true,
	}
}

// MarshalYAML marshals ARPValidate as YAML.
// This method used on marshaling YAML internally.
func (av *ARPValidate) MarshalYAML() (interface{}, error) {
	if av.isAssigned {
		return av.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals ARPValidate as YAML.
// This method used on unmarshaling YAML internally.
func (av *ARPValidate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val arpValidate
	if err := unmarshal(&val); err != nil {
		return err
	}

	av.val = val
	av.isAssigned = true

	return nil
}
