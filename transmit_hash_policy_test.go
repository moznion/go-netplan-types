package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyTransmitHashPolicy(t *testing.T) {
	given := TransmitHashPolicy{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal TransmitHashPolicy
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeTransmitHashPolicy(t *testing.T) {
	for _, given := range []*TransmitHashPolicy{
		Layer2TransmitHashPolicy(),
		Layer3And4TransmitHashPolicy(),
		Layer2And3TransmitHashPolicy(),
		Encap2And3TransmitHashPolicy(),
		Encap3And4TransmitHashPolicy(),
	} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal TransmitHashPolicy
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
