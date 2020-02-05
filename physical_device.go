package netplan

type PhysicalDevice struct {
	Match     *Match         `yaml:"match"`
	SetName   *NilableString `yaml:"set-name"`
	WakeOnLAN *NilableBool   `yaml:"wakeonlan"`
}
