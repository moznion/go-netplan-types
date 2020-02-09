package netplan

type PhysicalDevice struct {
	Match     *Match          `yaml:"match,omitempty"`
	SetName   *NillableString `yaml:"set-name,omitempty"`
	WakeOnLAN *NillableBool   `yaml:"wakeonlan,omitempty"`
}
