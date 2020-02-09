package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyRoutingPolicy(t *testing.T) {
	given := RoutingPolicy{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal RoutingPolicy
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeRoutingPolicy(t *testing.T) {
	given := RoutingPolicy{
		From: &Address{
			Address:   "192.0.2.1",
			PrefixLen: go_netplan_types.NilableUint8Of(32),
		},
		To: &Address{
			Address:   "0.0.0.0",
			PrefixLen: go_netplan_types.NilableUint8Of(0),
		},
		Table:         go_netplan_types.NilableUint64Of(200),
		Priority:      go_netplan_types.NilableUint32Of(100),
		Mark:          go_netplan_types.NilableUint64Of(1),
		TypeOfService: go_netplan_types.NilableUint8Of(8),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`from: 192.0.2.1/32
to: 0.0.0.0/0
table: 200
priority: 100
mark: 1
type-of-service: 8
`), marshal)

	var unmarshal RoutingPolicy
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
