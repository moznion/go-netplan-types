package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type BridgeParameters struct {
	AgeingTime   *go_netplan_types.NilableUint64 `yaml:"ageing-time,omitempty"`
	Priority     *go_netplan_types.NilableUint32 `yaml:"priority,omitempty"`
	PortPriority *go_netplan_types.NilableUint8  `yaml:"port-priority,omitempty"`
	ForwardDelay *go_netplan_types.NilableUint64 `yaml:"forward-delay,omitempty"`
	HelloTime    *go_netplan_types.NilableUint64 `yaml:"hello-time,omitempty"`
	MaxAge       *go_netplan_types.NilableUint64 `yaml:"max-age,omitempty"`
	PathCost     *go_netplan_types.NilableUint64 `yaml:"path-cost,omitempty"`
	STP          *go_netplan_types.NilableBool   `yaml:"stp,omitempty"`
}
