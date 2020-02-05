package netplan

type Match struct {
	Name       *NilableString `yaml:"name"`
	MacAddress *NilableString `yaml:"macaddress"`
	Driver     *NilableString `yaml:"driver"`
}
