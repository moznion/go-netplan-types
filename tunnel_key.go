package netplan

type TunnelKey struct {
	Input  *NilableUint64 `yaml:"input,omitempty"`
	Output *NilableUint64 `yaml:"output,omitempty"`
}
