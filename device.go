package netplan

type Device struct {
	DHCP4             *NilableBool   `yaml:"dhcp4,omitempty"`
	DHCP6             *NilableBool   `yaml:"dhcp6,omitempty"`
	IPv6Privacy       *NilableBool   `yaml:"ipv6-privacy,omitempty"`
	LinkLocal         []LinkLocal    `yaml:"link-local,omitempty"`
	Critical          *NilableBool   `yaml:"critical,omitempty"`
	DHCPIdentifier    *NilableString `yaml:"dhcp-identifier,omitempty"`
	DHCP4Overrides    *DHCPOverride  `yaml:"dhcp4-overrides,omitempty"`
	DHCP6Overrides    *DHCPOverride  `yaml:"dhcp6-overrides,omitempty"`
	AcceptRA          *NilableBool   `yaml:"accept-ra,omitempty"`
	Addresses         []*Address     `yaml:"addresses,omitempty"`
	Gateway4          *NilableString `yaml:"gateway4,omitempty"`
	Gateway6          *NilableString `yaml:"gateway6,omitempty"`
	NameServers       *Nameservers   `yaml:"nameservers,omitempty"`
	MacAddress        *NilableString `yaml:"macaddress,omitempty"`
	MTU               *NilableUint64 `yaml:"mtu,omitempty"`
	Optional          *NilableBool   `yaml:"optional,omitempty"`
	OptionalAddresses []string       `yaml:"optional-addresses,omitempty"`
	Routing           `yaml:"-,inline"`
}
