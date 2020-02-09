package netplan

// Route represents netplan's route configuration.
// See also: https://netplan.io/reference#routing
type Route struct {
	From   *Address        `yaml:"from,omitempty"`
	To     *Address        `yaml:"to,omitempty"`
	Via    *Address        `yaml:"via,omitempty"`
	OnLink *NillableBool   `yaml:"on-link,omitempty"`
	Metric *NillableUint64 `yaml:"metric,omitempty"`
	Type   *RouteType      `yaml:"type,omitempty"`
	Scope  *RouteScope     `yaml:"scope,omitempty"`
	Table  *NillableUint64 `yaml:"table,omitempty"`
}
