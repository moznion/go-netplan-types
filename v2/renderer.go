package netplan

type Renderer struct {
	val        renderer
	isAssigned bool
}

type renderer string

const (
	networkdRenderer       renderer = "networkd"
	networkManagerRenderer          = "NetworkManager"
)

func NetworkdRenderer() *Renderer {
	return &Renderer{
		val:        networkdRenderer,
		isAssigned: true,
	}
}

func NetworkManagerRenderer() *Renderer {
	return &Renderer{
		val:        networkManagerRenderer,
		isAssigned: true,
	}
}

func (r *Renderer) MarshalYAML() (interface{}, error) {
	if r.isAssigned {
		return r.val, nil
	}
	return nil, nil
}

func (r *Renderer) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val renderer
	if err := unmarshal(&val); err != nil {
		return err
	}

	r.val = val
	r.isAssigned = true

	return nil
}
