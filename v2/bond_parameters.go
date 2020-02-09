package netplan

type BondParameters struct {
	Mode                  *BondMode              `yaml:"mode,omitempty"`
	LACPRate              *LACPRate              `yaml:"lacp-rate,omitempty"`
	MIIMonitorInterval    *NillableUint64        `yaml:"mii-monitor-interval,omitempty"`
	MinLinks              *NillableUint64        `yaml:"min-links,omitempty"`
	TransmitHashPolicy    *TransmitHashPolicy    `yaml:"transmit-hash-policy,omitempty"`
	AdSelect              *AdSelect              `yaml:"ad-select,omitempty"`
	AllSlavesActive       *NillableBool          `yaml:"all-slaves-active,omitempty"`
	ARPInterval           *NillableUint64        `yaml:"arp-interval,omitempty"`
	ARPIPTargets          []string               `yaml:"arp-ip-targets,omitempty"`
	ARPValidate           *ARPValidate           `yaml:"arp-validate,omitempty"`
	ARPAllTargets         *ARPAllTargets         `yaml:"arp-all-targets,omitempty"`
	UpDelay               *NillableUint64        `yaml:"up-delay,omitempty"`
	DownDelay             *NillableUint64        `yaml:"down-delay,omitempty"`
	FailOverMACPolicy     *FailOverMACPolicy     `yaml:"fail-over-mac-policy,omitempty"`
	GratuitousARP         *NillableUint8         `yaml:"gratuitous-arp,omitempty"`
	PacketsPerSlave       *NillableUint16        `yaml:"packets-per-slave,omitempty"`
	PrimaryReselectPolicy *PrimaryReselectPolicy `yaml:"primary-reselect-policy,omitempty"`
	ResendIGMP            *NillableUint8         `yaml:"resend-igmp,omitempty"`
	LearnPacketInterval   *NillableUint32        `yaml:"learn-packet-interval,omitempty"`
	Primary               *NillableString        `yaml:"primary,omitempty"`
}
