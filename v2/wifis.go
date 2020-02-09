package netplan

// Wifi represents netplan's Wifi configuration.
// See also: https://netplan.io/reference#properties-for-device-type-wifis
type Wifi struct {
	Device         `yaml:"-,inline"`
	PhysicalDevice `yaml:"-,inline"`
	AccessPoints   AccessPoints    `yaml:"access-points,omitempty"`
	Auth           *Authentication `yaml:"auth,omitempty"`
}

// Wifis is a map that points Wifi interface name to Wifi configuration.
type Wifis map[string]*Wifi
