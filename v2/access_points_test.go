package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyAccessPoint(t *testing.T) {
	given := AccessPoint{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal AccessPoint
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeAccessPoint(t *testing.T) {
	given := AccessPoint{
		Password: yamlnillable.StringOf("pswd"),
		Mode:     InfrastructureAccessPointMode(),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`password: pswd
mode: infrastructure
`), marshal)

	var unmarshal AccessPoint
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
