package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyRouteScope(t *testing.T) {
	given := RouteScope{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal RouteScope
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeRouteScope(t *testing.T) {
	for _, given := range []*RouteScope{GlobalRouteScope(), LinkRouteScope(), HostRouteScope()} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal RouteScope
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
