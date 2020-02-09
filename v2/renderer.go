package netplan

// Renderer represents netplan's `renderer` as nillable.
type Renderer struct {
	val        renderer
	isAssigned bool
}

type renderer string

const (
	networkdRenderer       renderer = "networkd"
	networkManagerRenderer          = "NetworkManager"
)

// NetworkdRenderer returns `networkd` renderer.
func NetworkdRenderer() *Renderer {
	return &Renderer{
		val:        networkdRenderer,
		isAssigned: true,
	}
}

// NetworkManagerRenderer returns `NetworkManager` renderer.
func NetworkManagerRenderer() *Renderer {
	return &Renderer{
		val:        networkManagerRenderer,
		isAssigned: true,
	}
}

// MarshalYAML marshals Renderer as YAML.
// This method used on marshaling YAML internally.
func (r *Renderer) MarshalYAML() (interface{}, error) {
	if r.isAssigned {
		return r.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals Renderer as YAML.
// This method used on unmarshaling YAML internally.
func (r *Renderer) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val renderer
	if err := unmarshal(&val); err != nil {
		return err
	}

	r.val = val
	r.isAssigned = true

	return nil
}
