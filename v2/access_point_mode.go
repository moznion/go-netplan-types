package netplan

type AccessPointMode struct {
	val        accessPointMode
	isAssigned bool
}

type accessPointMode string

const (
	infrastructureAccessPointMode accessPointMode = "infrastructure"
	apAccessPointMode                             = "ap"
	adhocAccessPointMode                          = "adhoc"
)

func InfrastructureAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        infrastructureAccessPointMode,
		isAssigned: true,
	}
}

func APAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        apAccessPointMode,
		isAssigned: true,
	}
}

func AdhocAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        adhocAccessPointMode,
		isAssigned: true,
	}
}

func (apm *AccessPointMode) MarshalYAML() (interface{}, error) {
	if apm.isAssigned {
		return apm.val, nil
	}
	return nil, nil
}

func (apm *AccessPointMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val accessPointMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	apm.val = val
	apm.isAssigned = true

	return nil
}
