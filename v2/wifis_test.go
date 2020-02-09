package netplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyWifi(t *testing.T) {
	given := Wifi{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Wifi
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeWifi(t *testing.T) {
	given := Wifi{
		Device: Device{
			DHCP4: NilableBoolOf(true),
			DHCP6: NilableBoolOf(false),
		},
		PhysicalDevice: PhysicalDevice{
			WakeOnLAN: NilableBoolOf(false),
		},
		AccessPoints: map[string]*AccessPoint{
			"opennetwork": {},
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
wakeonlan: false
access-points:
  opennetwork: {}
auth:
  key-management: psk
  method: tls
`), marshal)

	var unmarshal Wifi
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
