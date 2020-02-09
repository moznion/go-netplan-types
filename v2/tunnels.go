package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Tunnel struct {
	Device `yaml:"-,inline"`
	Mode   *TunnelMode                      `yaml:"mode,omitempty"`
	Local  *go_netplan_types.NillableString `yaml:"local,omitempty"`
	Remote *go_netplan_types.NillableString `yaml:"remote,omitempty"`
	Key    *TunnelKey                       `yaml:"key,omitempty"`
	Keys   *TunnelKey                       `yaml:"keys,omitempty"`
}

type Tunnels map[string]*Tunnel
