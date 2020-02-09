package netplan

// VLAN represents netplan's VLAN configuration.
// See also: https://netplan.io/reference#properties-for-device-type-vlans
type VLAN struct {
	Device `yaml:"-,inline"`
	ID     *NillableUint16 `yaml:"id,omitempty"`
	Link   *NillableString `yaml:"link,omitempty"`
}

// VLANs is a map that points VLAN interface name to VLAN configuration.
type VLANs map[string]*VLAN
