package netplan

type RoutingPolicy struct {
	From          *Address       `yaml:"from,omitempty"`
	To            *Address       `yaml:"to,omitempty"`
	Table         *NilableUint64 `yaml:"table,omitempty"`
	Priority      *NilableUint32 `yaml:"priority,omitempty"`
	Mark          *NilableUint64 `yaml:"mark,omitempty"`
	TypeOfService *NilableUint8  `yaml:"type-of-service,omitempty"`
}
