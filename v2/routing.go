package netplan

// Routing represents netplan's routing configuration.
// See also: https://netplan.io/reference#routing
type Routing struct {
	Routes        []*Route         `yaml:"routes,omitempty"`
	RoutingPolicy []*RoutingPolicy `yaml:"routing-policy,omitempty"`
}
