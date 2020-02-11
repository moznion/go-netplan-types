package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// Tunnel represents netplan's tunnel configuration.
// See also: https://netplan.io/reference#properties-for-device-type-tunnels
type Tunnel struct {
	Device `yaml:"-,inline"`
	Mode   *TunnelMode          `yaml:"mode,omitempty"`
	Local  *yamlnillable.String `yaml:"local,omitempty"`
	Remote *yamlnillable.String `yaml:"remote,omitempty"`
	Key    *TunnelKey           `yaml:"key,omitempty"`
	Keys   *TunnelKey           `yaml:"keys,omitempty"`
}

// Tunnels is a map that points tunnel interface name to tunnel configuration.
type Tunnels map[string]*Tunnel
