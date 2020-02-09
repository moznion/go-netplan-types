package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type PhysicalDevice struct {
	Match     *Match                          `yaml:"match,omitempty"`
	SetName   *go_netplan_types.NilableString `yaml:"set-name,omitempty"`
	WakeOnLAN *go_netplan_types.NilableBool   `yaml:"wakeonlan,omitempty"`
}
