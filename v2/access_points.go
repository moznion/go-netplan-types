package netplan

type AccessPoint struct {
	Password *NillableString  `yaml:"password,omitempty"`
	Mode     *AccessPointMode `yaml:"mode,omitempty"`
}

type AccessPoints map[string]*AccessPoint
