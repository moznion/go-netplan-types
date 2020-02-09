package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyRouteType(t *testing.T) {
	given := RouteType{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal RouteType
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeRouteType(t *testing.T) {
	for _, given := range []*RouteType{UnicastRouteType(), UnreachableRouteType(), BlackholeRouteType(), ProhibitRouteType()} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal RouteType
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
