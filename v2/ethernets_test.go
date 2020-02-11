package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyEthernet(t *testing.T) {
	given := Ethernet{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Ethernet
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeEthernet(t *testing.T) {
	given := Ethernet{
		Device: Device{
			DHCP4: yamlnillable.BoolOf(true),
			DHCP6: yamlnillable.BoolOf(false),
		},
		PhysicalDevice: PhysicalDevice{
			Match: &Match{
				Name: yamlnillable.StringOf("dev-1"),
			},
		},
		Auth: &Authentication{
			KeyManagement: PSKKeyManagement(),
			Method:        TLSAuthMethod(),
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
match:
  name: dev-1
auth:
  key-management: psk
  method: tls
`), marshal)

	var unmarshal Ethernet
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
