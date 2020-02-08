package netplan

type Bridge struct {
	Device     `yaml:"-,inline,omitempty"`
	Interfaces []string          `yaml:"interfaces,omitempty"`
	Parameters *BridgeParameters `yaml:"parameters,omitempty"`
}

type Bridges map[string]*Bridge
