package example

import (
	"fmt"

	"github.com/moznion/go-netplan-types/v2"
	"gopkg.in/yaml.v2"
)

func Example() {
	network := netplan.Network{
		Network: &netplan.NetworkConfig{
			Ethernets: netplan.Ethernets{
				"eno1": &netplan.Ethernet{
					Device: netplan.Device{
						DHCP4: netplan.NillableBoolOf(true),
					},
				},
			},
		},
	}
	yamlString, _ := yaml.Marshal(&network)
	fmt.Printf("%s\n", yamlString)
	// This should be the following:
	// network:
	//   version: 2
	//   ethernets:
	//     eno1:
	//       dhcp4: true
}
