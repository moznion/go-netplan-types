package netplan

// Nameservers represents netplan's nameserver attribute for devices.
type Nameservers struct {
	Search    []string `yaml:"search,omitempty"`
	Addresses []string `yaml:"addresses,omitempty"`
}
