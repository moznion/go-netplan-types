package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
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
			DHCP4: go_netplan_types.NilableBoolOf(true),
			DHCP6: go_netplan_types.NilableBoolOf(false),
		},
		Mode:   GRETunnelMode(),
		Local:  go_netplan_types.NilableStringOf("192.0.2.1"),
		Remote: go_netplan_types.NilableStringOf("192.0.2.2"),
		Key: &TunnelKey{
			Input:  go_netplan_types.NilableUint64Of(1),
			Output: go_netplan_types.NilableUint64Of(2),
		},
		Keys: &TunnelKey{
			Input:  go_netplan_types.NilableUint64Of(3),
			Output: go_netplan_types.NilableUint64Of(4),
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
