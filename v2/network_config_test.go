package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyNetworkConfig(t *testing.T) {
	given := NetworkConfig{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`version: 2
`), marshal)

	var unmarshal NetworkConfig
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeNetworkConfig(t *testing.T) {
	given := NetworkConfig{
		Renderer: NetworkdRenderer(),
		Ethernets: Ethernets{
			"eth0": &Ethernet{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
		Wifis: Wifis{
			"wlan0": &Wifi{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
		Bridges: Bridges{
			"bridge0": &Bridge{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
		Bonds: Bonds{
			"bond0": &Bond{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
		Tunnels: Tunnels{
			"tunnel0": &Tunnel{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
		VLANs: VLANs{
			"vlan0": &VLAN{
				Device: Device{
					DHCP4: go_netplan_types.NilableBoolOf(true),
					DHCP6: go_netplan_types.NilableBoolOf(false),
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`version: 2
renderer: networkd
ethernets:
  eth0:
    dhcp4: true
    dhcp6: false
wifis:
  wlan0:
    dhcp4: true
    dhcp6: false
bridges:
  bridge0:
    dhcp4: true
    dhcp6: false
bonds:
  bond0:
    dhcp4: true
    dhcp6: false
tunnels:
  tunnel0:
    dhcp4: true
    dhcp6: false
vlans:
  vlan0:
    dhcp4: true
    dhcp6: false
`), marshal)

	var unmarshal NetworkConfig
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
