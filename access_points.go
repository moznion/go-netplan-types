package netplan

type AccessPoint struct {
	Password *NilableString   `yaml:"password,omitempty"`
	Mode     *AccessPointMode `yaml:"mode,omitempty"`
}

type AccessPoints map[string]*AccessPoint
