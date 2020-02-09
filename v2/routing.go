package netplan

type Routing struct {
	Routes        []*Route         `yaml:"routes,omitempty"`
	RoutingPolicy []*RoutingPolicy `yaml:"routing-policy,omitempty"`
}
