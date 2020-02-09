package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type RoutingPolicy struct {
	From          *Address                        `yaml:"from,omitempty"`
	To            *Address                        `yaml:"to,omitempty"`
	Table         *go_netplan_types.NilableUint64 `yaml:"table,omitempty"`
	Priority      *go_netplan_types.NilableUint32 `yaml:"priority,omitempty"`
	Mark          *go_netplan_types.NilableUint64 `yaml:"mark,omitempty"`
	TypeOfService *go_netplan_types.NilableUint8  `yaml:"type-of-service,omitempty"`
}
