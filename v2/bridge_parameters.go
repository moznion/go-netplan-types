package netplan

type BridgeParameters struct {
	AgeingTime   *NilableUint64 `yaml:"ageing-time,omitempty"`
	Priority     *NilableUint32 `yaml:"priority,omitempty"`
	PortPriority *NilableUint8  `yaml:"port-priority,omitempty"`
	ForwardDelay *NilableUint64 `yaml:"forward-delay,omitempty"`
	HelloTime    *NilableUint64 `yaml:"hello-time,omitempty"`
	MaxAge       *NilableUint64 `yaml:"max-age,omitempty"`
	PathCost     *NilableUint64 `yaml:"path-cost,omitempty"`
	STP          *NilableBool   `yaml:"stp,omitempty"`
}
