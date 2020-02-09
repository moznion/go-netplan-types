package netplan

// Bridge represents netplan's bridge interface.
// See also: https://netplan.io/reference#properties-for-device-type-bridges
type Bridge struct {
	Device     `yaml:"-,inline"`
	Interfaces []string          `yaml:"interfaces,omitempty"`
	Parameters *BridgeParameters `yaml:"parameters,omitempty"`
}

// Bridges is a map that points bridge interface name to bridge configuration.
type Bridges map[string]*Bridge
