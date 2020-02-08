package netplan

type Wifi struct {
	Device         `yaml:"-,inline,omitempty"`
	PhysicalDevice `yaml:"-,inline,omitempty"`
	AccessPoints   AccessPoints    `yaml:"access-points,omitempty"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

type Wifis map[string]*Wifi
