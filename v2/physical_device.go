package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// PhysicalDevice represents netplan's common properties for physical device types.
// See also: https://netplan.io/reference#common-properties-for-physical-device-types
type PhysicalDevice struct {
	Match     *Match               `yaml:"match,omitempty"`
	SetName   *yamlnillable.String `yaml:"set-name,omitempty"`
	WakeOnLAN *yamlnillable.Bool   `yaml:"wakeonlan,omitempty"`
}
