package netplan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyTunnelMode(t *testing.T) {
	given := TunnelMode{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`null
`), marshal)

	var unmarshal TunnelMode
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeTunnelMode(t *testing.T) {
	for _, given := range []*TunnelMode{
		SITTunnelMode(),
		GRETunnelMode(),
		IP6GRETunnelMode(),
		IPIPTunnelMode(),
		IPIP6TunnelMode(),
		IP6IP6TunnelMode(),
		VTITunnelMode(),
		VTI6TunnelMode(),
		GRETAPTunnelMode(),
		IP6GRETAPTunnelMode(),
		ISATAPTunnelMode(),
	} {
		marshal, err := yaml.Marshal(given)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(fmt.Sprintf(`%s
`, given.val)), marshal)

		var unmarshal TunnelMode
		err = yaml.Unmarshal(marshal, &unmarshal)
		assert.NoError(t, err)
		assert.EqualValues(t, given, &unmarshal)
	}
}
