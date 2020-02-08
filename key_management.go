package netplan

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

func NoneKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        noneKeyManagement,
		isAssigned: true,
	}
}

func PSKKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        pskKeyManagement,
		isAssigned: true,
	}
}

func EAPKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        eapKeyManagement,
		isAssigned: true,
	}
}

func IEEE8021xKeyManagement() *KeyManagement {
	return &KeyManagement{
		val:        ieee8021xKeyManagement,
		isAssigned: true,
	}
}

func (km *KeyManagement) MarshalYAML() (interface{}, error) {
	if km.isAssigned {
		return km.val, nil
	}
	return nil, nil
}

func (km *KeyManagement) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val keyManagement
	if err := unmarshal(&val); err != nil {
		return err
	}

	km.val = val
	km.isAssigned = true

	return nil
}
