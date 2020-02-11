package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// DHCPOverride represents netplan's DHCP overrides.
// See also: https://netplan.io/reference#dhcp-overrides
type DHCPOverride struct {
	UseDNS       *yamlnillable.Bool   `yaml:"use-dns,omitempty"`
	UseNTP       *yamlnillable.Bool   `yaml:"use-ntp,omitempty"`
	SendHostname *yamlnillable.Bool   `yaml:"send-hostname,omitempty"`
	UseHostname  *yamlnillable.Bool   `yaml:"use-hostname,omitempty"`
	UseMTU       *yamlnillable.Bool   `yaml:"use-mtu,omitempty"`
	Hostname     *yamlnillable.String `yaml:"hostname,omitempty"`
	UseRoutes    *yamlnillable.Bool   `yaml:"use-routes,omitempty"`
	RouteMetric  *yamlnillable.Uint64 `yaml:"route-metric,omitempty"`
}
