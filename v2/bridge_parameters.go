package netplan

type BridgeParameters struct {
	AgeingTime   *NillableUint64 `yaml:"ageing-time,omitempty"`
	Priority     *NillableUint32 `yaml:"priority,omitempty"`
	PortPriority *NillableUint8  `yaml:"port-priority,omitempty"`
	ForwardDelay *NillableUint64 `yaml:"forward-delay,omitempty"`
	HelloTime    *NillableUint64 `yaml:"hello-time,omitempty"`
	MaxAge       *NillableUint64 `yaml:"max-age,omitempty"`
	PathCost     *NillableUint64 `yaml:"path-cost,omitempty"`
	STP          *NillableBool   `yaml:"stp,omitempty"`
}
