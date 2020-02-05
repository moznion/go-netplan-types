package netplan

type DHCPOverride struct {
	UseDNS       *NilableBool   `yaml:"use-dns"`
	UseNTP       *NilableBool   `yaml:"use-ntp"`
	SendHostname *NilableBool   `yaml:"send-hostname"`
	UseHostname  *NilableBool   `yaml:"use-hostname"`
	UseMTU       *NilableBool   `yaml:"use-mtu"`
	Hostname     *NilableString `yaml:"hostname"`
	UseRoutes    *NilableBool   `yaml:"use-routes"`
	RouteMetric  *NilableUint64 `yaml:"route-metric"`
}
