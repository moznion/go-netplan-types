package netplan

// RoutingPolicy represents netplan's routing policy configuration.
// See also: https://netplan.io/reference#routing
type RoutingPolicy struct {
	From          *Address        `yaml:"from,omitempty"`
	To            *Address        `yaml:"to,omitempty"`
	Table         *NillableUint64 `yaml:"table,omitempty"`
	Priority      *NillableUint32 `yaml:"priority,omitempty"`
	Mark          *NillableUint64 `yaml:"mark,omitempty"`
	TypeOfService *NillableUint8  `yaml:"type-of-service,omitempty"`
}
