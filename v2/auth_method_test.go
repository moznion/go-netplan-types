package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyAuthMethod(t *testing.T) {
	given := AuthMethod{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal AuthMethod
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeAuthMethod(t *testing.T) {
	for _, given := range []*AuthMethod{TLSAuthMethod(), PEAPAuthMethod(), TTLSAuthMethod()} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal AuthMethod
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
