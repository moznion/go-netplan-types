package netplan

const networkVersion2 = 2

// NetworkConfig represents netplan's network config.
// See also: https://netplan.io/reference#properties-for-device-type-ethernets
type NetworkConfig struct {
	VersionMustBe2 uint8     `yaml:"version"`
	Renderer       *Renderer `yaml:"renderer,omitempty"`
	Ethernets      Ethernets `yaml:"ethernets,omitempty"`
	Wifis          Wifis     `yaml:"wifis,omitempty"`
	Bridges        Bridges   `yaml:"bridges,omitempty"`
	Bonds          Bonds     `yaml:"bonds,omitempty"`
	Tunnels        Tunnels   `yaml:"tunnels,omitempty"`
	VLANs          VLANs     `yaml:"vlans,omitempty"`
}

type NetworkConfigForMarshalling struct {
	VersionMustBe2 uint8     `yaml:"version"`
	Renderer       *Renderer `yaml:"renderer,omitempty"`
	Ethernets      Ethernets `yaml:"ethernets,omitempty"`
	Wifis          Wifis     `yaml:"wifis,omitempty"`
	Bridges        Bridges   `yaml:"bridges,omitempty"`
	Bonds          Bonds     `yaml:"bonds,omitempty"`
	Tunnels        Tunnels   `yaml:"tunnels,omitempty"`
	VLANs          VLANs     `yaml:"vlans,omitempty"`
}

type networkConfigForUnmarshalling struct {
	Renderer  *Renderer `yaml:"renderer,omitempty"`
	Ethernets Ethernets `yaml:"ethernets,omitempty"`
	Wifis     Wifis     `yaml:"wifis,omitempty"`
	Bridges   Bridges   `yaml:"bridges,omitempty"`
	Bonds     Bonds     `yaml:"bonds,omitempty"`
	Tunnels   Tunnels   `yaml:"tunnels,omitempty"`
	VLANs     VLANs     `yaml:"vlans,omitempty"`
}

// MarshalYAML marshals NetworkConfig as YAML.
// This method used on marshaling YAML internally.
func (n *NetworkConfig) MarshalYAML() (interface{}, error) {
	nm := NetworkConfigForMarshalling{
		VersionMustBe2: networkVersion2,
		Renderer:       n.Renderer,
		Ethernets:      n.Ethernets,
		Wifis:          n.Wifis,
		Bridges:        n.Bridges,
		Bonds:          n.Bonds,
		Tunnels:        n.Tunnels,
		VLANs:          n.VLANs,
	}

	return nm, nil
}

// UnmarshalYAML unmarshals NetworkConfig as YAML.
// This method used on unmarshaling YAML internally.
func (n *NetworkConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var network networkConfigForUnmarshalling
	if err := unmarshal(&network); err != nil {
		return err
	}

	n.VersionMustBe2 = networkVersion2
	n.Renderer = network.Renderer
	n.Ethernets = network.Ethernets
	n.Wifis = network.Wifis
	n.Bridges = network.Bridges
	n.Bonds = network.Bonds
	n.Tunnels = network.Tunnels
	n.VLANs = network.VLANs

	return nil
}
