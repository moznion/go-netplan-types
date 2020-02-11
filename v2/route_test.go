package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyRoute(t *testing.T) {
	given := Route{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Route
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeRoute(t *testing.T) {
	given := Route{
		From: &Address{
			Address:   "192.0.2.1",
			PrefixLen: yamlnillable.Uint8Of(32),
		},
		To: &Address{
			Address:   "0.0.0.0",
			PrefixLen: yamlnillable.Uint8Of(0),
		},
		Via: &Address{
			Address: "192.0.2.2",
		},
		OnLink: yamlnillable.BoolOf(false),
		Metric: yamlnillable.Uint64Of(100),
		Type:   UnicastRouteType(),
		Scope:  LinkRouteScope(),
		Table:  yamlnillable.Uint64Of(200),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`from: 192.0.2.1/32
to: 0.0.0.0/0
via: 192.0.2.2
on-link: false
metric: 100
type: unicast
scope: link
table: 200
`), marshal)

	var unmarshal Route
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
