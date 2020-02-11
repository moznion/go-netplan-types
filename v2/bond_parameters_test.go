package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyBondParameters(t *testing.T) {
	given := BondParameters{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal BondParameters
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeBondParameters(t *testing.T) {
	given := BondParameters{
		Mode:                  BalanceRRBondMode(),
		LACPRate:              FastLACPRate(),
		MIIMonitorInterval:    yamlnillable.Uint64Of(1),
		MinLinks:              yamlnillable.Uint64Of(2),
		TransmitHashPolicy:    Layer2TransmitHashPolicy(),
		AdSelect:              BandwidthAdSelect(),
		AllSlavesActive:       yamlnillable.BoolOf(false),
		ARPInterval:           yamlnillable.Uint64Of(3),
		ARPIPTargets:          []string{"foo"},
		ARPValidate:           ActiveArpValidate(),
		ARPAllTargets:         AnyARPAllTargets(),
		UpDelay:               yamlnillable.Uint64Of(4),
		DownDelay:             yamlnillable.Uint64Of(5),
		FailOverMACPolicy:     NoneFailOverMACPolicy(),
		GratuitousARP:         yamlnillable.Uint8Of(6),
		PacketsPerSlave:       yamlnillable.Uint16Of(7),
		PrimaryReselectPolicy: AlwaysPrimaryReselectPolicy(),
		ResendIGMP:            yamlnillable.Uint8Of(8),
		LearnPacketInterval:   yamlnillable.Uint32Of(9),
		Primary:               yamlnillable.StringOf("bar"),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`mode: balance-rr
lacp-rate: fast
mii-monitor-interval: 1
min-links: 2
transmit-hash-policy: layer2
ad-select: bandwidth
all-slaves-active: false
arp-interval: 3
arp-ip-targets:
- foo
arp-validate: active
arp-all-targets: any
up-delay: 4
down-delay: 5
fail-over-mac-policy: none
gratuitous-arp: 6
packets-per-slave: 7
primary-reselect-policy: always
resend-igmp: 8
learn-packet-interval: 9
primary: bar
`), marshal)

	var unmarshal BondParameters
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
