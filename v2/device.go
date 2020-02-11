package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// Device represents netplan's common properties for all device types.
// See also: https://netplan.io/reference#common-properties-for-all-device-types
type Device struct {
	DHCP4             *yamlnillable.Bool   `yaml:"dhcp4,omitempty"`
	DHCP6             *yamlnillable.Bool   `yaml:"dhcp6,omitempty"`
	IPv6Privacy       *yamlnillable.Bool   `yaml:"ipv6-privacy,omitempty"`
	LinkLocal         []LinkLocal          `yaml:"link-local,omitempty"`
	Critical          *yamlnillable.Bool   `yaml:"critical,omitempty"`
	DHCPIdentifier    *yamlnillable.String `yaml:"dhcp-identifier,omitempty"`
	DHCP4Overrides    *DHCPOverride        `yaml:"dhcp4-overrides,omitempty"`
	DHCP6Overrides    *DHCPOverride        `yaml:"dhcp6-overrides,omitempty"`
	AcceptRA          *yamlnillable.Bool   `yaml:"accept-ra,omitempty"`
	Addresses         []*Address           `yaml:"addresses,omitempty"`
	Gateway4          *yamlnillable.String `yaml:"gateway4,omitempty"`
	Gateway6          *yamlnillable.String `yaml:"gateway6,omitempty"`
	NameServers       *Nameservers         `yaml:"nameservers,omitempty"`
	MacAddress        *yamlnillable.String `yaml:"macaddress,omitempty"`
	MTU               *yamlnillable.Uint64 `yaml:"mtu,omitempty"`
	Optional          *yamlnillable.Bool   `yaml:"optional,omitempty"`
	OptionalAddresses []string             `yaml:"optional-addresses,omitempty"`
	Renderer          *Renderer            `yaml:"renderer,omitempty"`
	Routing           `yaml:"-,inline"`
}
