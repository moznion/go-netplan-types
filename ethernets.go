package netplan

type Ethernet struct {
	Device         `yaml:"-,inline"`
	PhysicalDevice `yaml:"-,inline"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

type Ethernets map[string]*Ethernet
