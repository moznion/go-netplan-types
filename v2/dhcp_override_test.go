package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyDHCPOverride(t *testing.T) {
	given := DHCPOverride{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal DHCPOverride
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeDHCPOverride(t *testing.T) {
	given := DHCPOverride{
		UseDNS:       yamlnillable.BoolOf(false),
		UseNTP:       yamlnillable.BoolOf(true),
		SendHostname: yamlnillable.BoolOf(false),
		UseHostname:  yamlnillable.BoolOf(true),
		UseMTU:       yamlnillable.BoolOf(false),
		Hostname:     yamlnillable.StringOf("host"),
		UseRoutes:    yamlnillable.BoolOf(true),
		RouteMetric:  yamlnillable.Uint64Of(100),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`use-dns: false
use-ntp: true
send-hostname: false
use-hostname: true
use-mtu: false
hostname: host
use-routes: true
route-metric: 100
`), marshal)

	var unmarshal DHCPOverride
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
