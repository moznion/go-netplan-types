package netplan

type Ethernet struct {
	Device         `yaml:"-,inline,omitempty"`
	PhysicalDevice `yaml:"-,inline,omitempty"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

type Ethernets map[string]*Ethernet
