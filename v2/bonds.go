package netplan

// Bond represents netplan's bond interface.
// See also: https://netplan.io/reference#properties-for-device-type-bonds
type Bond struct {
	Device     `yaml:"-,inline"`
	Parameters *BondParameters `yaml:"parameters,omitempty"`
	Interfaces []string        `yaml:"interfaces,omitempty"`
}

// Bonds is a map that points bond interface name to bond configuration.
type Bonds map[string]*Bond
