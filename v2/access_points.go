package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// AccessPoint represents netplan's access point.
type AccessPoint struct {
	Password *yamlnillable.String `yaml:"password,omitempty"`
	Mode     *AccessPointMode     `yaml:"mode,omitempty"`
}

// AccessPoints is a map that points access-point name to access-point configuration.
type AccessPoints map[string]*AccessPoint
