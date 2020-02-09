package netplan

type Tunnel struct {
	Device `yaml:"-,inline"`
	Mode   *TunnelMode     `yaml:"mode,omitempty"`
	Local  *NillableString `yaml:"local,omitempty"`
	Remote *NillableString `yaml:"remote,omitempty"`
	Key    *TunnelKey      `yaml:"key,omitempty"`
	Keys   *TunnelKey      `yaml:"keys,omitempty"`
}

type Tunnels map[string]*Tunnel
