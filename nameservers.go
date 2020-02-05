package netplan

type Nameservers struct {
	Search    []string `yaml:"search"`
	Addresses []string `yaml:"addresses"`
}
