package netplan

import (
	"testing"

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
		AgeingTime:   NillableUint64Of(1),
		Priority:     NillableUint32Of(2),
		PortPriority: NillableUint8Of(3),
		ForwardDelay: NillableUint64Of(4),
		HelloTime:    NillableUint64Of(5),
		MaxAge:       NillableUint64Of(6),
		PathCost:     NillableUint64Of(7),
		STP:          NillableBoolOf(false),
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
