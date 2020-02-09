package netplan

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
