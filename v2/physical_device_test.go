package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
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
			Name:       yamlnillable.StringOf("name"),
			MacAddress: nil,
			Driver:     nil,
		},
		SetName:   yamlnillable.StringOf("setname"),
		WakeOnLAN: yamlnillable.BoolOf(false),
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
