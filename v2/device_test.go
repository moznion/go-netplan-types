package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyDevice(t *testing.T) {
	given := Device{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Device
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeDevice(t *testing.T) {
	given := Device{
		DHCP4:          yamlnillable.BoolOf(true),
		DHCP6:          yamlnillable.BoolOf(false),
		IPv6Privacy:    yamlnillable.BoolOf(false),
		LinkLocal:      []LinkLocal{IPv4LinkLocal},
		Critical:       yamlnillable.BoolOf(true),
		DHCPIdentifier: yamlnillable.StringOf("mac"),
		DHCP4Overrides: &DHCPOverride{
			UseDNS:       yamlnillable.BoolOf(true),
			UseNTP:       yamlnillable.BoolOf(true),
			SendHostname: yamlnillable.BoolOf(true),
			UseHostname:  yamlnillable.BoolOf(true),
			UseMTU:       yamlnillable.BoolOf(true),
			Hostname:     yamlnillable.StringOf("host"),
			UseRoutes:    yamlnillable.BoolOf(true),
			RouteMetric:  yamlnillable.Uint64Of(100),
		},
		DHCP6Overrides: nil,
		AcceptRA:       yamlnillable.BoolOf(false),
		Addresses: []*Address{
			{
				Address:   "192.0.2.1",
				PrefixLen: yamlnillable.Uint8Of(32),
			},
			{
				Address:   "192.0.2.2",
				PrefixLen: yamlnillable.Uint8Of(32),
			},
		},
		Gateway4: yamlnillable.StringOf("192.0.2.254"),
		Gateway6: nil,
		NameServers: &Nameservers{
			Search:    []string{"domain-1", "domain-2"},
			Addresses: []string{"8.8.8.8"},
		},
		MacAddress:        yamlnillable.StringOf("de:ad:be:ef:ca:fe"),
		MTU:               yamlnillable.Uint64Of(1500),
		Optional:          yamlnillable.BoolOf(false),
		OptionalAddresses: nil,
		Routing: Routing{
			Routes: []*Route{
				{
					From: &Address{
						Address:   "198.0.2.1",
						PrefixLen: yamlnillable.Uint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: yamlnillable.Uint8Of(0),
					},
				},
				{
					From: &Address{
						Address:   "198.0.2.2",
						PrefixLen: yamlnillable.Uint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: yamlnillable.Uint8Of(0),
					},
				},
			},
			RoutingPolicy: []*RoutingPolicy{
				{
					From: &Address{
						Address:   "192.0.2.0",
						PrefixLen: yamlnillable.Uint8Of(24),
					},
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
ipv6-privacy: false
link-local:
- ipv4
critical: true
dhcp-identifier: mac
dhcp4-overrides:
  use-dns: true
  use-ntp: true
  send-hostname: true
  use-hostname: true
  use-mtu: true
  hostname: host
  use-routes: true
  route-metric: 100
accept-ra: false
addresses:
- 192.0.2.1/32
- 192.0.2.2/32
gateway4: 192.0.2.254
nameservers:
  search:
  - domain-1
  - domain-2
  addresses:
  - 8.8.8.8
macaddress: de:ad:be:ef:ca:fe
mtu: 1500
optional: false
routes:
- from: 198.0.2.1/32
  to: 0.0.0.0/0
- from: 198.0.2.2/32
  to: 0.0.0.0/0
routing-policy:
- from: 192.0.2.0/24
`), marshal)

	var unmarshal Device
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
