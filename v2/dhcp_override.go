package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type DHCPOverride struct {
	UseDNS       *go_netplan_types.NillableBool   `yaml:"use-dns,omitempty"`
	UseNTP       *go_netplan_types.NillableBool   `yaml:"use-ntp,omitempty"`
	SendHostname *go_netplan_types.NillableBool   `yaml:"send-hostname,omitempty"`
	UseHostname  *go_netplan_types.NillableBool   `yaml:"use-hostname,omitempty"`
	UseMTU       *go_netplan_types.NillableBool   `yaml:"use-mtu,omitempty"`
	Hostname     *go_netplan_types.NillableString `yaml:"hostname,omitempty"`
	UseRoutes    *go_netplan_types.NillableBool   `yaml:"use-routes,omitempty"`
	RouteMetric  *go_netplan_types.NillableUint64 `yaml:"route-metric,omitempty"`
}
