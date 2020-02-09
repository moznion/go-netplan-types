package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type DHCPOverride struct {
	UseDNS       *go_netplan_types.NilableBool   `yaml:"use-dns,omitempty"`
	UseNTP       *go_netplan_types.NilableBool   `yaml:"use-ntp,omitempty"`
	SendHostname *go_netplan_types.NilableBool   `yaml:"send-hostname,omitempty"`
	UseHostname  *go_netplan_types.NilableBool   `yaml:"use-hostname,omitempty"`
	UseMTU       *go_netplan_types.NilableBool   `yaml:"use-mtu,omitempty"`
	Hostname     *go_netplan_types.NilableString `yaml:"hostname,omitempty"`
	UseRoutes    *go_netplan_types.NilableBool   `yaml:"use-routes,omitempty"`
	RouteMetric  *go_netplan_types.NilableUint64 `yaml:"route-metric,omitempty"`
}
