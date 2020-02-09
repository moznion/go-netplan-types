package netplan

type Match struct {
	Name       *NillableString `yaml:"name,omitempty"`
	MacAddress *NillableString `yaml:"macaddress,omitempty"`
	Driver     *NillableString `yaml:"driver,omitempty"`
}
