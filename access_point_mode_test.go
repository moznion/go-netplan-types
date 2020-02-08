package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyAccessPointMode(t *testing.T) {
	given := AccessPointMode{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal AccessPointMode
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeAccessPointMode(t *testing.T) {
	for _, given := range []*AccessPointMode{InfrastructureAccessPointMode(), APAccessPointMode(), AdhocAccessPointMode()} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal AccessPointMode
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
