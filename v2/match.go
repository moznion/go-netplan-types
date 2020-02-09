package netplan

// Match represents netplan's match attribute for devices.
type Match struct {
	Name       *NillableString `yaml:"name,omitempty"`
	MacAddress *NillableString `yaml:"macaddress,omitempty"`
	Driver     *NillableString `yaml:"driver,omitempty"`
}
