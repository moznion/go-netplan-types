package netplan

// AccessPointMode represents netplan's access point mode as nillable.
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

// InfrastructureAccessPointMode returns `infrastructure` access point mode.
func InfrastructureAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        infrastructureAccessPointMode,
		isAssigned: true,
	}
}

// APAccessPointMode returns `ap` access point mode.
func APAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        apAccessPointMode,
		isAssigned: true,
	}
}

// AdhocAccessPointMode returns `adhoc` access point mode.
func AdhocAccessPointMode() *AccessPointMode {
	return &AccessPointMode{
		val:        adhocAccessPointMode,
		isAssigned: true,
	}
}

// MarshalYAML marshals AccessPointMode as YAML.
// This method used on marshaling YAML internally.
func (apm *AccessPointMode) MarshalYAML() (interface{}, error) {
	if apm.isAssigned {
		return apm.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals AccessPointMode as YAML.
// This method used on unmarshaling YAML internally.
func (apm *AccessPointMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val accessPointMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	apm.val = val
	apm.isAssigned = true

	return nil
}
