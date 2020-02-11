package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyTunnel(t *testing.T) {
	given := Tunnel{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Tunnel
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeTunnel(t *testing.T) {
	given := Tunnel{
		Device: Device{
			DHCP4: yamlnillable.BoolOf(true),
			DHCP6: yamlnillable.BoolOf(false),
		},
		Mode:   GRETunnelMode(),
		Local:  yamlnillable.StringOf("192.0.2.1"),
		Remote: yamlnillable.StringOf("192.0.2.2"),
		Key: &TunnelKey{
			Input:  yamlnillable.Uint64Of(1),
			Output: yamlnillable.Uint64Of(2),
		},
		Keys: &TunnelKey{
			Input:  yamlnillable.Uint64Of(3),
			Output: yamlnillable.Uint64Of(4),
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
mode: gre
local: 192.0.2.1
remote: 192.0.2.2
key:
  input: 1
  output: 2
keys:
  input: 3
  output: 4
`), marshal)

	var unmarshal Tunnel
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
