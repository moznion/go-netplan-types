package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// Route represents netplan's route configuration.
// See also: https://netplan.io/reference#routing
type Route struct {
	From   *Address             `yaml:"from,omitempty"`
	To     *Address             `yaml:"to,omitempty"`
	Via    *Address             `yaml:"via,omitempty"`
	OnLink *yamlnillable.Bool   `yaml:"on-link,omitempty"`
	Metric *yamlnillable.Uint64 `yaml:"metric,omitempty"`
	Type   *RouteType           `yaml:"type,omitempty"`
	Scope  *RouteScope          `yaml:"scope,omitempty"`
	Table  *yamlnillable.Uint64 `yaml:"table,omitempty"`
}
