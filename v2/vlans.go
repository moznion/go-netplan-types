package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type VLAN struct {
	Device `yaml:"-,inline"`
	ID     *go_netplan_types.NilableUint16 `yaml:"id,omitempty"`
	Link   *go_netplan_types.NilableString `yaml:"link,omitempty"`
}

type VLANs map[string]*VLAN
