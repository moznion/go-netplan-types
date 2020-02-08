package netplan

type Route struct {
	From   *Address       `yaml:"from,omitempty"`
	To     *Address       `yaml:"to,omitempty"`
	Via    *Address       `yaml:"via,omitempty"`
	OnLink *NilableBool   `yaml:"on-link,omitempty"`
	Metric *NilableUint64 `yaml:"metric,omitempty"`
	Type   *RouteType     `yaml:"type,omitempty"`
	Scope  *RouteScope    `yaml:"scope,omitempty"`
	Table  *NilableUint64 `yaml:"table,omitempty"`
}
