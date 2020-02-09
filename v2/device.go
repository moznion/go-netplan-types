package netplan

type Device struct {
	DHCP4             *NillableBool   `yaml:"dhcp4,omitempty"`
	DHCP6             *NillableBool   `yaml:"dhcp6,omitempty"`
	IPv6Privacy       *NillableBool   `yaml:"ipv6-privacy,omitempty"`
	LinkLocal         []LinkLocal     `yaml:"link-local,omitempty"`
	Critical          *NillableBool   `yaml:"critical,omitempty"`
	DHCPIdentifier    *NillableString `yaml:"dhcp-identifier,omitempty"`
	DHCP4Overrides    *DHCPOverride   `yaml:"dhcp4-overrides,omitempty"`
	DHCP6Overrides    *DHCPOverride   `yaml:"dhcp6-overrides,omitempty"`
	AcceptRA          *NillableBool   `yaml:"accept-ra,omitempty"`
	Addresses         []*Address      `yaml:"addresses,omitempty"`
	Gateway4          *NillableString `yaml:"gateway4,omitempty"`
	Gateway6          *NillableString `yaml:"gateway6,omitempty"`
	NameServers       *Nameservers    `yaml:"nameservers,omitempty"`
	MacAddress        *NillableString `yaml:"macaddress,omitempty"`
	MTU               *NillableUint64 `yaml:"mtu,omitempty"`
	Optional          *NillableBool   `yaml:"optional,omitempty"`
	OptionalAddresses []string        `yaml:"optional-addresses,omitempty"`
	Renderer          *Renderer       `yaml:"renderer,omitempty"`
	Routing           `yaml:"-,inline"`
}
