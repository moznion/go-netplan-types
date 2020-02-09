package netplan

type AuthMethod struct {
	val        authMethod
	isAssigned bool
}

type authMethod string

const (
	tlsAuthMethod  authMethod = "tls"
	peapAuthMethod            = "peap"
	ttlsAuthMethod            = "ttls"
)

func TLSAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        tlsAuthMethod,
		isAssigned: true,
	}
}

func PEAPAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        peapAuthMethod,
		isAssigned: true,
	}
}

func TTLSAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        ttlsAuthMethod,
		isAssigned: true,
	}
}

func (am *AuthMethod) MarshalYAML() (interface{}, error) {
	if am.isAssigned {
		return am.val, nil
	}
	return nil, nil
}

func (am *AuthMethod) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val authMethod
	if err := unmarshal(&val); err != nil {
		return err
	}

	am.val = val
	am.isAssigned = true

	return nil
}
