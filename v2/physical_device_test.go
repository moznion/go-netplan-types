package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyPhysicalDevice(t *testing.T) {
	given := PhysicalDevice{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal PhysicalDevice
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializePhysicalDevice(t *testing.T) {
	given := PhysicalDevice{
		Match: &Match{
			Name:       go_netplan_types.NilableStringOf("name"),
			MacAddress: nil,
			Driver:     nil,
		},
		SetName:   go_netplan_types.NilableStringOf("setname"),
		WakeOnLAN: go_netplan_types.NilableBoolOf(false),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`match:
  name: name
set-name: setname
wakeonlan: false
`), marshal)

	var unmarshal PhysicalDevice
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
