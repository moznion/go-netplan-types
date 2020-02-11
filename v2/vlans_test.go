package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyVLAN(t *testing.T) {
	given := VLAN{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal VLAN
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeVLAN(t *testing.T) {
	given := VLAN{
		Device: Device{
			DHCP4: yamlnillable.BoolOf(true),
			DHCP6: yamlnillable.BoolOf(false),
		},
		ID:   yamlnillable.Uint16Of(1),
		Link: yamlnillable.StringOf("link-1"),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
id: 1
link: link-1
`), marshal)

	var unmarshal VLAN
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
