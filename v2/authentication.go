package netplan

import yamlnillable "github.com/moznion/go-yaml-nillable"

// Authentication represents netplan's authentication attribute.
// See also: https://netplan.io/reference#authentication
type Authentication struct {
	KeyManagement     *KeyManagement       `yaml:"key-management,omitempty"`
	Password          *yamlnillable.String `yaml:"password,omitempty"`
	Method            *AuthMethod          `yaml:"method,omitempty"`
	Identity          *yamlnillable.String `yaml:"identity,omitempty"`
	AnonymousIdentity *yamlnillable.String `yaml:"anonymous-identity,omitempty"`
	CACertificate     *yamlnillable.String `yaml:"ca-certificate,omitempty"`
	ClientCertificate *yamlnillable.String `yaml:"client-certificate,omitempty"`
	ClientKey         *yamlnillable.String `yaml:"client-key,omitempty"`
	ClientKeyPassword *yamlnillable.String `yaml:"client-key-password,omitempty"`
}
