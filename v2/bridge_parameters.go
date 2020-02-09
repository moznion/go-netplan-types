package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type BridgeParameters struct {
	AgeingTime   *go_netplan_types.NillableUint64 `yaml:"ageing-time,omitempty"`
	Priority     *go_netplan_types.NillableUint32 `yaml:"priority,omitempty"`
	PortPriority *go_netplan_types.NillableUint8  `yaml:"port-priority,omitempty"`
	ForwardDelay *go_netplan_types.NillableUint64 `yaml:"forward-delay,omitempty"`
	HelloTime    *go_netplan_types.NillableUint64 `yaml:"hello-time,omitempty"`
	MaxAge       *go_netplan_types.NillableUint64 `yaml:"max-age,omitempty"`
	PathCost     *go_netplan_types.NillableUint64 `yaml:"path-cost,omitempty"`
	STP          *go_netplan_types.NillableBool   `yaml:"stp,omitempty"`
}
