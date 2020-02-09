package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyBridgeParameters(t *testing.T) {
	given := BridgeParameters{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal BridgeParameters
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeBridgeParameters(t *testing.T) {
	given := BridgeParameters{
		AgeingTime:   go_netplan_types.NilableUint64Of(1),
		Priority:     go_netplan_types.NilableUint32Of(2),
		PortPriority: go_netplan_types.NilableUint8Of(3),
		ForwardDelay: go_netplan_types.NilableUint64Of(4),
		HelloTime:    go_netplan_types.NilableUint64Of(5),
		MaxAge:       go_netplan_types.NilableUint64Of(6),
		PathCost:     go_netplan_types.NilableUint64Of(7),
		STP:          go_netplan_types.NilableBoolOf(false),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`ageing-time: 1
priority: 2
port-priority: 3
forward-delay: 4
hello-time: 5
max-age: 6
path-cost: 7
stp: false
`), marshal)

	var unmarshal BridgeParameters
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
