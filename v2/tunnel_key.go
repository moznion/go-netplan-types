package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// TunnelKey represents netplan's tunnel key.
type TunnelKey struct {
	Input  *yamlnillable.Uint64 `yaml:"input,omitempty"`
	Output *yamlnillable.Uint64 `yaml:"output,omitempty"`
}
