package netplan

// DHCPOverride represents netplan's DHCP overrides.
// See also: https://netplan.io/reference#dhcp-overrides
type DHCPOverride struct {
	UseDNS       *NillableBool   `yaml:"use-dns,omitempty"`
	UseNTP       *NillableBool   `yaml:"use-ntp,omitempty"`
	SendHostname *NillableBool   `yaml:"send-hostname,omitempty"`
	UseHostname  *NillableBool   `yaml:"use-hostname,omitempty"`
	UseMTU       *NillableBool   `yaml:"use-mtu,omitempty"`
	Hostname     *NillableString `yaml:"hostname,omitempty"`
	UseRoutes    *NillableBool   `yaml:"use-routes,omitempty"`
	RouteMetric  *NillableUint64 `yaml:"route-metric,omitempty"`
}
