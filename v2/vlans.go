package netplan

type VLAN struct {
	Device `yaml:"-,inline"`
	ID     *NilableUint16 `yaml:"id,omitempty"`
	Link   *NilableString `yaml:"link,omitempty"`
}

type VLANs map[string]*VLAN
