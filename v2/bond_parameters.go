package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type BondParameters struct {
	Mode                  *BondMode                        `yaml:"mode,omitempty"`
	LACPRate              *LACPRate                        `yaml:"lacp-rate,omitempty"`
	MIIMonitorInterval    *go_netplan_types.NillableUint64 `yaml:"mii-monitor-interval,omitempty"`
	MinLinks              *go_netplan_types.NillableUint64 `yaml:"min-links,omitempty"`
	TransmitHashPolicy    *TransmitHashPolicy              `yaml:"transmit-hash-policy,omitempty"`
	AdSelect              *AdSelect                        `yaml:"ad-select,omitempty"`
	AllSlavesActive       *go_netplan_types.NillableBool   `yaml:"all-slaves-active,omitempty"`
	ARPInterval           *go_netplan_types.NillableUint64 `yaml:"arp-interval,omitempty"`
	ARPIPTargets          []string                         `yaml:"arp-ip-targets,omitempty"`
	ARPValidate           *ARPValidate                     `yaml:"arp-validate,omitempty"`
	ARPAllTargets         *ARPAllTargets                   `yaml:"arp-all-targets,omitempty"`
	UpDelay               *go_netplan_types.NillableUint64 `yaml:"up-delay,omitempty"`
	DownDelay             *go_netplan_types.NillableUint64 `yaml:"down-delay,omitempty"`
	FailOverMACPolicy     *FailOverMACPolicy               `yaml:"fail-over-mac-policy,omitempty"`
	GratuitousARP         *go_netplan_types.NillableUint8  `yaml:"gratuitous-arp,omitempty"`
	PacketsPerSlave       *go_netplan_types.NillableUint16 `yaml:"packets-per-slave,omitempty"`
	PrimaryReselectPolicy *PrimaryReselectPolicy           `yaml:"primary-reselect-policy,omitempty"`
	ResendIGMP            *go_netplan_types.NillableUint8  `yaml:"resend-igmp,omitempty"`
	LearnPacketInterval   *go_netplan_types.NillableUint32 `yaml:"learn-packet-interval,omitempty"`
	Primary               *go_netplan_types.NillableString `yaml:"primary,omitempty"`
}
