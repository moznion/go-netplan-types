package netplan

// AuthMethod represents netplan's authentication method as nillable.
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

// TLSAuthMethod returns `tls` authentication method.
func TLSAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        tlsAuthMethod,
		isAssigned: true,
	}
}

// PEAPAuthMethod returns `peap` authentication method.
func PEAPAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        peapAuthMethod,
		isAssigned: true,
	}
}

// TTLSAuthMethod returns `ttls` authentication method.
func TTLSAuthMethod() *AuthMethod {
	return &AuthMethod{
		val:        ttlsAuthMethod,
		isAssigned: true,
	}
}

// MarshalYAML marshals AuthMethod as YAML.
// This method used on marshaling YAML internally.
func (am *AuthMethod) MarshalYAML() (interface{}, error) {
	if am.isAssigned {
		return am.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals AuthMethod as YAML.
// This method used on unmarshaling YAML internally.
func (am *AuthMethod) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val authMethod
	if err := unmarshal(&val); err != nil {
		return err
	}

	am.val = val
	am.isAssigned = true

	return nil
}
