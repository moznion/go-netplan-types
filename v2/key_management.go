package netplan

// KeyManagement represents netplan's kind of key-management as nillable.
type KeyManagement struct {
	val        keyManagement
	isAssigned bool
}

type keyManagement string

const (
	noneKeyManagement      keyManagement = "none"
	pskKeyManagement                     = "psk"
	eapKeyManagement                     = "eap"
	ieee8021xKeyManagement               = "802.1x"
)

// NoneKeyManagement returns `none` key-management type.
func NoneKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        noneKeyManagement,
		isAssigned: true,
	}
}

// PSKKeyManagement returns `psk` key-management type.
func PSKKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        pskKeyManagement,
		isAssigned: true,
	}
}

// EAPKeyManagement returns `eap` key-management type.
func EAPKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        eapKeyManagement,
		isAssigned: true,
	}
}

// IEEE8021xKeyManagement returns `802.1x` key-management type.
func IEEE8021xKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        ieee8021xKeyManagement,
		isAssigned: true,
	}
}

// MarshalYAML marshals KeyManagement as YAML.
// This method used on marshaling YAML internally.
func (km *KeyManagement) MarshalYAML() (interface{}, error) {
	if km.isAssigned {
		return km.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals KeyManagement as YAML.
// This method used on unmarshaling YAML internally.
func (km *KeyManagement) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val keyManagement
	if err := unmarshal(&val); err != nil {
		return err
	}

	km.val = val
	km.isAssigned = true

	return nil
}
