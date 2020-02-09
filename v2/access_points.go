package netplan

// AccessPoint represents netplan's access point.
type AccessPoint struct {
	Password *NillableString  `yaml:"password,omitempty"`
	Mode     *AccessPointMode `yaml:"mode,omitempty"`
}

// AccessPoints is a map that points access-point name to access-point configuration.
type AccessPoints map[string]*AccessPoint
