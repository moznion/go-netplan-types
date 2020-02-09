package netplan

// Authentication represents netplan's authentication attribute.
// See also: https://netplan.io/reference#authentication
type Authentication struct {
	KeyManagement     *KeyManagement  `yaml:"key-management,omitempty"`
	Password          *NillableString `yaml:"password,omitempty"`
	Method            *AuthMethod     `yaml:"method,omitempty"`
	Identity          *NillableString `yaml:"identity,omitempty"`
	AnonymousIdentity *NillableString `yaml:"anonymous-identity,omitempty"`
	CACertificate     *NillableString `yaml:"ca-certificate,omitempty"`
	ClientCertificate *NillableString `yaml:"client-certificate,omitempty"`
	ClientKey         *NillableString `yaml:"client-key,omitempty"`
	ClientKeyPassword *NillableString `yaml:"client-key-password,omitempty"`
}
