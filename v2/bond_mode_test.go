package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyBondMode(t *testing.T) {
	given := BondMode{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal BondMode
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeBondMode(t *testing.T) {
	for _, given := range []*BondMode{
		BalanceRRBondMode(),
		ActiveBackupBondMode(),
		BalanceXORBondMode(),
		BroadcastBondMode(),
		IEEE8023adBondMode(),
		BalanceTLBBondMode(),
		BalanceALBBondMode(),
	} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal BondMode
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
