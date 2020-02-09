package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
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
		DHCP4:          go_netplan_types.NillableBoolOf(true),
		DHCP6:          go_netplan_types.NillableBoolOf(false),
		IPv6Privacy:    go_netplan_types.NillableBoolOf(false),
		LinkLocal:      []LinkLocal{IPv4LinkLocal},
		Critical:       go_netplan_types.NillableBoolOf(true),
		DHCPIdentifier: go_netplan_types.NillableStringOf("mac"),
		DHCP4Overrides: &DHCPOverride{
			UseDNS:       go_netplan_types.NillableBoolOf(true),
			UseNTP:       go_netplan_types.NillableBoolOf(true),
			SendHostname: go_netplan_types.NillableBoolOf(true),
			UseHostname:  go_netplan_types.NillableBoolOf(true),
			UseMTU:       go_netplan_types.NillableBoolOf(true),
			Hostname:     go_netplan_types.NillableStringOf("host"),
			UseRoutes:    go_netplan_types.NillableBoolOf(true),
			RouteMetric:  go_netplan_types.NillableUint64Of(100),
		},
		DHCP6Overrides: nil,
		AcceptRA:       go_netplan_types.NillableBoolOf(false),
		Addresses: []*Address{
			{
				Address:   "192.0.2.1",
				PrefixLen: go_netplan_types.NillableUint8Of(32),
			},
			{
				Address:   "192.0.2.2",
				PrefixLen: go_netplan_types.NillableUint8Of(32),
			},
		},
		Gateway4: go_netplan_types.NillableStringOf("192.0.2.254"),
		Gateway6: nil,
		NameServers: &Nameservers{
			Search:    []string{"domain-1", "domain-2"},
			Addresses: []string{"8.8.8.8"},
		},
		MacAddress:        go_netplan_types.NillableStringOf("de:ad:be:ef:ca:fe"),
		MTU:               go_netplan_types.NillableUint64Of(1500),
		Optional:          go_netplan_types.NillableBoolOf(false),
		OptionalAddresses: nil,
		Routing: Routing{
			Routes: []*Route{
				{
					From: &Address{
						Address:   "198.0.2.1",
						PrefixLen: go_netplan_types.NillableUint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: go_netplan_types.NillableUint8Of(0),
					},
				},
				{
					From: &Address{
						Address:   "198.0.2.2",
						PrefixLen: go_netplan_types.NillableUint8Of(32),
					},
					To: &Address{
						Address:   "0.0.0.0",
						PrefixLen: go_netplan_types.NillableUint8Of(0),
					},
				},
			},
			RoutingPolicy: []*RoutingPolicy{
				{
					From: &Address{
						Address:   "192.0.2.0",
						PrefixLen: go_netplan_types.NillableUint8Of(24),
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
