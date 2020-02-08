package netplan

type Bond struct {
	Device     `yaml:"-,inline"`
	Parameters *BondParameters `yaml:"parameters,omitempty"`
	Interfaces []string        `yaml:"interfaces,omitempty"`
}

type Bonds map[string]*Bond
