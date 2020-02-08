package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyPrimaryReselectPolicy(t *testing.T) {
	given := PrimaryReselectPolicy{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal PrimaryReselectPolicy
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializePrimaryReselectPolicy(t *testing.T) {
	for _, given := range []*PrimaryReselectPolicy{AlwaysPrimaryReselectPolicy(), BetterPrimaryReselectPolicy(), FailurePrimaryReselectPolicy()} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal PrimaryReselectPolicy
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
