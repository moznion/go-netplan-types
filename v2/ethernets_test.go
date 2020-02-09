package netplan

import (
	"testing"

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
			DHCP4: NilableBoolOf(true),
			DHCP6: NilableBoolOf(false),
		},
		PhysicalDevice: PhysicalDevice{
			Match: &Match{
				Name: NilableStringOf("dev-1"),
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
