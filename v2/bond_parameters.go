package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// BondParameters represents netplan's bond parameters.
// See also: https://netplan.io/reference#properties-for-device-type-bonds
type BondParameters struct {
	Mode                  *BondMode              `yaml:"mode,omitempty"`
	LACPRate              *LACPRate              `yaml:"lacp-rate,omitempty"`
	MIIMonitorInterval    *yamlnillable.Uint64   `yaml:"mii-monitor-interval,omitempty"`
	MinLinks              *yamlnillable.Uint64   `yaml:"min-links,omitempty"`
	TransmitHashPolicy    *TransmitHashPolicy    `yaml:"transmit-hash-policy,omitempty"`
	AdSelect              *AdSelect              `yaml:"ad-select,omitempty"`
	AllSlavesActive       *yamlnillable.Bool     `yaml:"all-slaves-active,omitempty"`
	ARPInterval           *yamlnillable.Uint64   `yaml:"arp-interval,omitempty"`
	ARPIPTargets          []string               `yaml:"arp-ip-targets,omitempty"`
	ARPValidate           *ARPValidate           `yaml:"arp-validate,omitempty"`
	ARPAllTargets         *ARPAllTargets         `yaml:"arp-all-targets,omitempty"`
	UpDelay               *yamlnillable.Uint64   `yaml:"up-delay,omitempty"`
	DownDelay             *yamlnillable.Uint64   `yaml:"down-delay,omitempty"`
	FailOverMACPolicy     *FailOverMACPolicy     `yaml:"fail-over-mac-policy,omitempty"`
	GratuitousARP         *yamlnillable.Uint8    `yaml:"gratuitous-arp,omitempty"`
	PacketsPerSlave       *yamlnillable.Uint16   `yaml:"packets-per-slave,omitempty"`
	PrimaryReselectPolicy *PrimaryReselectPolicy `yaml:"primary-reselect-policy,omitempty"`
	ResendIGMP            *yamlnillable.Uint8    `yaml:"resend-igmp,omitempty"`
	LearnPacketInterval   *yamlnillable.Uint32   `yaml:"learn-packet-interval,omitempty"`
	Primary               *yamlnillable.String   `yaml:"primary,omitempty"`
}
