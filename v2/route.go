package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Route struct {
	From   *Address                         `yaml:"from,omitempty"`
	To     *Address                         `yaml:"to,omitempty"`
	Via    *Address                         `yaml:"via,omitempty"`
	OnLink *go_netplan_types.NillableBool   `yaml:"on-link,omitempty"`
	Metric *go_netplan_types.NillableUint64 `yaml:"metric,omitempty"`
	Type   *RouteType                       `yaml:"type,omitempty"`
	Scope  *RouteScope                      `yaml:"scope,omitempty"`
	Table  *go_netplan_types.NillableUint64 `yaml:"table,omitempty"`
}
