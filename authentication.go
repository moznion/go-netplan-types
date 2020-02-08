package netplan

type Authentication struct {
	KeyManagement     *KeyManagement `yaml:"key-management,omitempty"`
	Password          *NilableString `yaml:"password,omitempty"`
	Method            *AuthMethod    `yaml:"method,omitempty"`
	Identity          *NilableString `yaml:"identity,omitempty"`
	AnonymousIdentity *NilableString `yaml:"anonymous-identity,omitempty"`
	CACertificate     *NilableString `yaml:"ca-certificate,omitempty"`
	ClientCertificate *NilableString `yaml:"client-certificate,omitempty"`
	ClientKey         *NilableString `yaml:"client-key,omitempty"`
	ClientKeyPassword *NilableString `yaml:"client-key-password,omitempty"`
}
