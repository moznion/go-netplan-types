package netplan

type Device struct {
	DHCP4             *NilableBool   `yaml:"dhcp4"`
	DHCP6             *NilableBool   `yaml:"dhcp6"`
	IPv6Privacy       *NilableBool   `yaml:"ipv6-privacy"`
	LinkLocal         []LinkLocal    `yaml:"link-local"`
	Critical          *NilableBool   `yaml:"critical"`
	DHCPIdentifier    *NilableString `yaml:"dhcp-identifier"`
	DHCP4Overrides    *DHCPOverride  `yaml:"dhcp4-overrides"`
	DHCP6Overrides    *DHCPOverride  `yaml:"dhcp6-overrides"`
	AcceptRA          *NilableBool   `yaml:"accept-ra"`
	Addresses         []*Address     `yaml:"addresses"`
	Gateway4          *NilableString `yaml:"gateway4"`
	Gateway6          *NilableString `yaml:"gateway6"`
	NameServers       *Nameservers   `yaml:"nameservers"`
	MacAddress        *NilableString `yaml:"macaddress"`
	MTU               *NilableUint64 `yaml:"mtu"`
	Optional          *NilableBool   `yaml:"optional"`
	OptionalAddresses []string       `yaml:"optional-addresses"`
	Routing           `yaml:"-,inline"`
}
