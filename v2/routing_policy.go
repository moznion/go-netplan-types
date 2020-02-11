package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// RoutingPolicy represents netplan's routing policy configuration.
// See also: https://netplan.io/reference#routing
type RoutingPolicy struct {
	From          *Address             `yaml:"from,omitempty"`
	To            *Address             `yaml:"to,omitempty"`
	Table         *yamlnillable.Uint64 `yaml:"table,omitempty"`
	Priority      *yamlnillable.Uint32 `yaml:"priority,omitempty"`
	Mark          *yamlnillable.Uint64 `yaml:"mark,omitempty"`
	TypeOfService *yamlnillable.Uint8  `yaml:"type-of-service,omitempty"`
}
