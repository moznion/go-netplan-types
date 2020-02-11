package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// Match represents netplan's match attribute for devices.
type Match struct {
	Name       *yamlnillable.String `yaml:"name,omitempty"`
	MacAddress *yamlnillable.String `yaml:"macaddress,omitempty"`
	Driver     *yamlnillable.String `yaml:"driver,omitempty"`
}
