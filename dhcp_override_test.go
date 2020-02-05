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
	assert.EqualValues(t, []byte(`use-dns: null
use-ntp: null
send-hostname: null
use-hostname: null
use-mtu: null
hostname: null
use-routes: null
route-metric: null
`), marshal)

	var unmarshal DHCPOverride
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeDHCPOverride(t *testing.T) {
	given := DHCPOverride{
		UseDNS:       NilableBoolOf(false),
		UseNTP:       NilableBoolOf(true),
		SendHostname: NilableBoolOf(false),
		UseHostname:  NilableBoolOf(true),
		UseMTU:       NilableBoolOf(false),
		Hostname:     NilableStringOf("host"),
		UseRoutes:    NilableBoolOf(true),
		RouteMetric:  NilableUint64Of(100),
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
