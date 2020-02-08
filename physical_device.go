package netplan

type PhysicalDevice struct {
	Match     *Match         `yaml:"match,omitempty"`
	SetName   *NilableString `yaml:"set-name,omitempty"`
	WakeOnLAN *NilableBool   `yaml:"wakeonlan,omitempty"`
}
