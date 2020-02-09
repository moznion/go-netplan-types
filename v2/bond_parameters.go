package netplan

type BondParameters struct {
	Mode                  *BondMode              `yaml:"mode,omitempty"`
	LACPRate              *LACPRate              `yaml:"lacp-rate,omitempty"`
	MIIMonitorInterval    *NilableUint64         `yaml:"mii-monitor-interval,omitempty"`
	MinLinks              *NilableUint64         `yaml:"min-links,omitempty"`
	TransmitHashPolicy    *TransmitHashPolicy    `yaml:"transmit-hash-policy,omitempty"`
	AdSelect              *AdSelect              `yaml:"ad-select,omitempty"`
	AllSlavesActive       *NilableBool           `yaml:"all-slaves-active,omitempty"`
	ARPInterval           *NilableUint64         `yaml:"arp-interval,omitempty"`
	ARPIPTargets          []string               `yaml:"arp-ip-targets,omitempty"`
	ARPValidate           *ARPValidate           `yaml:"arp-validate,omitempty"`
	ARPAllTargets         *ARPAllTargets         `yaml:"arp-all-targets,omitempty"`
	UpDelay               *NilableUint64         `yaml:"up-delay,omitempty"`
	DownDelay             *NilableUint64         `yaml:"down-delay,omitempty"`
	FailOverMACPolicy     *FailOverMACPolicy     `yaml:"fail-over-mac-policy,omitempty"`
	GratuitousARP         *NilableUint8          `yaml:"gratuitous-arp,omitempty"`
	PacketsPerSlave       *NilableUint16         `yaml:"packets-per-slave,omitempty"`
	PrimaryReselectPolicy *PrimaryReselectPolicy `yaml:"primary-reselect-policy,omitempty"`
	ResendIGMP            *NilableUint8          `yaml:"resend-igmp,omitempty"`
	LearnPacketInterval   *NilableUint32         `yaml:"learn-packet-interval,omitempty"`
	Primary               *NilableString         `yaml:"primary,omitempty"`
}
