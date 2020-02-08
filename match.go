package netplan

type Match struct {
	Name       *NilableString `yaml:"name,omitempty"`
	MacAddress *NilableString `yaml:"macaddress,omitempty"`
	Driver     *NilableString `yaml:"driver,omitempty"`
}
