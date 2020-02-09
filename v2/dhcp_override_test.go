package netplan

import (
	"testing"

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
		UseDNS:       NillableBoolOf(false),
		UseNTP:       NillableBoolOf(true),
		SendHostname: NillableBoolOf(false),
		UseHostname:  NillableBoolOf(true),
		UseMTU:       NillableBoolOf(false),
		Hostname:     NillableStringOf("host"),
		UseRoutes:    NillableBoolOf(true),
		RouteMetric:  NillableUint64Of(100),
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
