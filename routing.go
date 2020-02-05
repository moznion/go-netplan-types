package netplan

type Routing struct {
	Routes        []*Route         `yaml:"routes"`
	RoutingPolicy []*RoutingPolicy `yaml:"routing-policy"`
}
