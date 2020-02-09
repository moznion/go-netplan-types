package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"

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
		UseDNS:       go_netplan_types.NilableBoolOf(false),
		UseNTP:       go_netplan_types.NilableBoolOf(true),
		SendHostname: go_netplan_types.NilableBoolOf(false),
		UseHostname:  go_netplan_types.NilableBoolOf(true),
		UseMTU:       go_netplan_types.NilableBoolOf(false),
		Hostname:     go_netplan_types.NilableStringOf("host"),
		UseRoutes:    go_netplan_types.NilableBoolOf(true),
		RouteMetric:  go_netplan_types.NilableUint64Of(100),
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
