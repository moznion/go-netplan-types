package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Device struct {
	DHCP4             *go_netplan_types.NilableBool   `yaml:"dhcp4,omitempty"`
	DHCP6             *go_netplan_types.NilableBool   `yaml:"dhcp6,omitempty"`
	IPv6Privacy       *go_netplan_types.NilableBool   `yaml:"ipv6-privacy,omitempty"`
	LinkLocal         []LinkLocal                     `yaml:"link-local,omitempty"`
	Critical          *go_netplan_types.NilableBool   `yaml:"critical,omitempty"`
	DHCPIdentifier    *go_netplan_types.NilableString `yaml:"dhcp-identifier,omitempty"`
	DHCP4Overrides    *DHCPOverride                   `yaml:"dhcp4-overrides,omitempty"`
	DHCP6Overrides    *DHCPOverride                   `yaml:"dhcp6-overrides,omitempty"`
	AcceptRA          *go_netplan_types.NilableBool   `yaml:"accept-ra,omitempty"`
	Addresses         []*Address                      `yaml:"addresses,omitempty"`
	Gateway4          *go_netplan_types.NilableString `yaml:"gateway4,omitempty"`
	Gateway6          *go_netplan_types.NilableString `yaml:"gateway6,omitempty"`
	NameServers       *Nameservers                    `yaml:"nameservers,omitempty"`
	MacAddress        *go_netplan_types.NilableString `yaml:"macaddress,omitempty"`
	MTU               *go_netplan_types.NilableUint64 `yaml:"mtu,omitempty"`
	Optional          *go_netplan_types.NilableBool   `yaml:"optional,omitempty"`
	OptionalAddresses []string                        `yaml:"optional-addresses,omitempty"`
	Renderer          *Renderer                       `yaml:"renderer,omitempty"`
	Routing           `yaml:"-,inline"`
}
