package netplan

type TunnelKey struct {
	Input  *NillableUint64 `yaml:"input,omitempty"`
	Output *NillableUint64 `yaml:"output,omitempty"`
}
