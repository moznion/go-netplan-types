package netplan

type RoutingPolicy struct {
	From          *Address        `yaml:"from,omitempty"`
	To            *Address        `yaml:"to,omitempty"`
	Table         *NillableUint64 `yaml:"table,omitempty"`
	Priority      *NillableUint32 `yaml:"priority,omitempty"`
	Mark          *NillableUint64 `yaml:"mark,omitempty"`
	TypeOfService *NillableUint8  `yaml:"type-of-service,omitempty"`
}
