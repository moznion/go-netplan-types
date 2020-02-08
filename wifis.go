package netplan

type Wifi struct {
	Device         `yaml:"-,inline"`
	PhysicalDevice `yaml:"-,inline"`
	AccessPoints   AccessPoints    `yaml:"access-points,omitempty"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

type Wifis map[string]*Wifi
