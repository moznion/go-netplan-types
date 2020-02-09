package netplan

// Network represents the top level structure for netplan's configuration.
type Network struct {
	Network *NetworkConfig `yaml:"network"`
}
