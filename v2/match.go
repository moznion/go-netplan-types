package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Match struct {
	Name       *go_netplan_types.NilableString `yaml:"name,omitempty"`
	MacAddress *go_netplan_types.NilableString `yaml:"macaddress,omitempty"`
	Driver     *go_netplan_types.NilableString `yaml:"driver,omitempty"`
}
