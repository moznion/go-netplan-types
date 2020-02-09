package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Match struct {
	Name       *go_netplan_types.NillableString `yaml:"name,omitempty"`
	MacAddress *go_netplan_types.NillableString `yaml:"macaddress,omitempty"`
	Driver     *go_netplan_types.NillableString `yaml:"driver,omitempty"`
}
