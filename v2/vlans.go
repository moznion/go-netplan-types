package netplan

type VLAN struct {
	Device `yaml:"-,inline"`
	ID     *NillableUint16 `yaml:"id,omitempty"`
	Link   *NillableString `yaml:"link,omitempty"`
}

type VLANs map[string]*VLAN
