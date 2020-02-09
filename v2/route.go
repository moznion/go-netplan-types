package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Route struct {
	From   *Address                        `yaml:"from,omitempty"`
	To     *Address                        `yaml:"to,omitempty"`
	Via    *Address                        `yaml:"via,omitempty"`
	OnLink *go_netplan_types.NilableBool   `yaml:"on-link,omitempty"`
	Metric *go_netplan_types.NilableUint64 `yaml:"metric,omitempty"`
	Type   *RouteType                      `yaml:"type,omitempty"`
	Scope  *RouteScope                     `yaml:"scope,omitempty"`
	Table  *go_netplan_types.NilableUint64 `yaml:"table,omitempty"`
}
