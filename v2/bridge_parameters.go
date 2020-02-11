package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// BridgeParameters represents netplan's bridge parameters.
// See also: https://netplan.io/reference#properties-for-device-type-bridges
type BridgeParameters struct {
	AgeingTime   *yamlnillable.Uint64 `yaml:"ageing-time,omitempty"`
	Priority     *yamlnillable.Uint32 `yaml:"priority,omitempty"`
	PortPriority *yamlnillable.Uint8  `yaml:"port-priority,omitempty"`
	ForwardDelay *yamlnillable.Uint64 `yaml:"forward-delay,omitempty"`
	HelloTime    *yamlnillable.Uint64 `yaml:"hello-time,omitempty"`
	MaxAge       *yamlnillable.Uint64 `yaml:"max-age,omitempty"`
	PathCost     *yamlnillable.Uint64 `yaml:"path-cost,omitempty"`
	STP          *yamlnillable.Bool   `yaml:"stp,omitempty"`
}
