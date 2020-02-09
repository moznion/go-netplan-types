package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type AccessPoint struct {
	Password *go_netplan_types.NilableString `yaml:"password,omitempty"`
	Mode     *AccessPointMode                `yaml:"mode,omitempty"`
}

type AccessPoints map[string]*AccessPoint
