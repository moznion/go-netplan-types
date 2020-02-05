package netplan

type RoutingPolicy struct {
	From          *Address       `yaml:"from"`
	To            *Address       `yaml:"to"`
	Table         *NilableUint64 `yaml:"table"`
	Priority      *NilableUint32 `yaml:"priority"`
	Mark          *NilableUint64 `yaml:"mark"`
	TypeOfService *NilableUint8  `yaml:"type-of-service"`
}
