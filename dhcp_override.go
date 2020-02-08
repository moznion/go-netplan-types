package netplan

type DHCPOverride struct {
	UseDNS       *NilableBool   `yaml:"use-dns,omitempty"`
	UseNTP       *NilableBool   `yaml:"use-ntp,omitempty"`
	SendHostname *NilableBool   `yaml:"send-hostname,omitempty"`
	UseHostname  *NilableBool   `yaml:"use-hostname,omitempty"`
	UseMTU       *NilableBool   `yaml:"use-mtu,omitempty"`
	Hostname     *NilableString `yaml:"hostname,omitempty"`
	UseRoutes    *NilableBool   `yaml:"use-routes,omitempty"`
	RouteMetric  *NilableUint64 `yaml:"route-metric,omitempty"`
}
