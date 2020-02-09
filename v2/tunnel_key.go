package netplan

// TunnelKey represents netplan's tunnel key.
type TunnelKey struct {
	Input  *NillableUint64 `yaml:"input,omitempty"`
	Output *NillableUint64 `yaml:"output,omitempty"`
}
