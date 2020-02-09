package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type TunnelKey struct {
	Input  *go_netplan_types.NilableUint64 `yaml:"input,omitempty"`
	Output *go_netplan_types.NilableUint64 `yaml:"output,omitempty"`
}
