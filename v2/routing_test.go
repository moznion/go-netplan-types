package netplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyRouting(t *testing.T) {
	given := Routing{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Routing
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeRouting(t *testing.T) {
	given := Routing{
		Routes: []*Route{
			{
				From: &Address{
					Address:   "192.0.2.1",
					PrefixLen: NilableUint8Of(32),
				},
				To: &Address{
					Address:   "0.0.0.0",
					PrefixLen: NilableUint8Of(0),
				},
			},
			{
				From: &Address{
					Address:   "192.0.2.2",
					PrefixLen: NilableUint8Of(32),
				},
				To: &Address{
					Address:   "0.0.0.0",
					PrefixLen: NilableUint8Of(0),
				},
			},
		},
		RoutingPolicy: []*RoutingPolicy{
			{
				From: &Address{
					Address:   "192.0.2.100",
					PrefixLen: NilableUint8Of(0),
				},
				To: &Address{
					Address:   "0.0.0.0",
					PrefixLen: NilableUint8Of(0),
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`routes:
- from: 192.0.2.1/32
  to: 0.0.0.0/0
- from: 192.0.2.2/32
  to: 0.0.0.0/0
routing-policy:
- from: 192.0.2.100/0
  to: 0.0.0.0/0
`), marshal)

	var unmarshal Routing
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
