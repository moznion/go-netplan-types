package netplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyDevice(t *testing.T) {
	given := Device{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: null
dhcp6: null
ipv6-privacy: null
link-local: []
critical: null
dhcp-identifier: null
dhcp4-overrides: null
dhcp6-overrides: null
accept-ra: null
addresses: []
gateway4: null
gateway6: null
nameservers: null
macaddress: null
mtu: null
optional: null
optional-addresses: []
routes: []
routing-policy: []
`), marshal)

	var unmarshal Device
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, Device{
		LinkLocal:         []LinkLocal{},
		Addresses:         []*Address{},
		OptionalAddresses: []string{},
		Routing: Routing{
			Routes:        []*Route{},
			RoutingPolicy: []*RoutingPolicy{},
		},
	}, unmarshal)
}

func TestSerializeDevice(t *testing.T) {
	given := Device{
		DHCP4:          NilableBoolOf(true),
		DHCP6:          NilableBoolOf(false),
		IPv6Privacy:    NilableBoolOf(false),
		LinkLocal:      []LinkLocal{IPv4LinkLocal},
		Critical:       NilableBoolOf(true),
		DHCPIdentifier: NilableStringOf("mac"),
		DHCP4Overrides: &DHCPOverride{
			UseDNS:       NilableBoolOf(true),
			UseNTP:       NilableBoolOf(true),
			SendHostname: NilableBoolOf(true),
			UseHostname:  NilableBoolOf(true),
			UseMTU:       NilableBoolOf(true),
			Hostname:     NilableStringOf("host"),
			UseRoutes:    NilableBoolOf(true),
			RouteMetric:  NilableUint64Of(100),
		},
		DHCP6Overrides: nil,
		AcceptRA:       NilableBoolOf(false),
		Addresses: []*Address{
			{
				Address:   "192.0.2.1",
				PrefixLen: NilableUint8Of(32),
			},
			{
				Address:   "192.0.2.2",
				PrefixLen: NilableUint8Of(32),
			},
		},
		Gateway4: NilableStringOf("192.0.2.254"),
		Gateway6: nil,
		NameServers: &Nameservers{
			Search:    []string{"domain-1", "domain-2"},
			Addresses: []string{"8.8.8.8"},
		},
		MacAddress:        NilableStringOf("de:ad:be:ef:ca:fe"),
		MTU:               NilableUint64Of(1500),
		Optional:          NilableBoolOf(false),
		OptionalAddresses: nil,
		Routing: Routing{
			Routes: []*Route{
				{
					From: &Address{
						Address:   "198.0.2.1",
						PrefixLen: NilableUint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: NilableUint8Of(0),
					},
				},
				{
					From: &Address{
						Address:   "198.0.2.2",
						PrefixLen: NilableUint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: NilableUint8Of(0),
					},
				},
			},
			RoutingPolicy: []*RoutingPolicy{
				{
					From: &Address{
						Address:   "192.0.2.0",
						PrefixLen: NilableUint8Of(24),
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
dhcp6-overrides: null
accept-ra: false
addresses:
- 192.0.2.1/32
- 192.0.2.2/32
gateway4: 192.0.2.254
gateway6: null
nameservers:
  search:
  - domain-1
  - domain-2
  addresses:
  - 8.8.8.8
macaddress: de:ad:be:ef:ca:fe
mtu: 1500
optional: false
optional-addresses: []
routes:
- from: 198.0.2.1/32
  to: 0.0.0.0/0
  via: null
  on-link: null
  metric: null
  type: null
  scope: null
  table: null
- from: 198.0.2.2/32
  to: 0.0.0.0/0
  via: null
  on-link: null
  metric: null
  type: null
  scope: null
  table: null
routing-policy:
- from: 192.0.2.0/24
  to: null
  table: null
  priority: null
  mark: null
  type-of-service: null
`), marshal)

	var unmarshal Device
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)

	given.OptionalAddresses = []string{}
	assert.EqualValues(t, given, unmarshal)
}
