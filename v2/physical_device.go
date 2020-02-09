package netplan

// PhysicalDevice represents netplan's common properties for physical device types.
// See also: https://netplan.io/reference#common-properties-for-physical-device-types
type PhysicalDevice struct {
	Match     *Match          `yaml:"match,omitempty"`
	SetName   *NillableString `yaml:"set-name,omitempty"`
	WakeOnLAN *NillableBool   `yaml:"wakeonlan,omitempty"`
}
