package netplan

// Ethernet represents netplan's ethernet interface.
// See also: https://netplan.io/reference#properties-for-device-type-ethernets
type Ethernet struct {
	Device         `yaml:"-,inline"`
	PhysicalDevice `yaml:"-,inline"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

// Ethernets is a map that points ethernet interface name to ethernet configuration.
type Ethernets map[string]*Ethernet
