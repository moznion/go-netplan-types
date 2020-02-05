package netplan

type Route struct {
	From   *Address       `yaml:"from"`
	To     *Address       `yaml:"to"`
	Via    *Address       `yaml:"via"`
	OnLink *NilableBool   `yaml:"on-link"`
	Metric *NilableUint64 `yaml:"metric"`
	Type   *RouteType     `yaml:"type"`
	Scope  *RouteScope    `yaml:"scope"`
	Table  *NilableUint64 `yaml:"table"`
}
