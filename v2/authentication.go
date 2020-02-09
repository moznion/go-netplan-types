package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Authentication struct {
	KeyManagement     *KeyManagement                  `yaml:"key-management,omitempty"`
	Password          *go_netplan_types.NilableString `yaml:"password,omitempty"`
	Method            *AuthMethod                     `yaml:"method,omitempty"`
	Identity          *go_netplan_types.NilableString `yaml:"identity,omitempty"`
	AnonymousIdentity *go_netplan_types.NilableString `yaml:"anonymous-identity,omitempty"`
	CACertificate     *go_netplan_types.NilableString `yaml:"ca-certificate,omitempty"`
	ClientCertificate *go_netplan_types.NilableString `yaml:"client-certificate,omitempty"`
	ClientKey         *go_netplan_types.NilableString `yaml:"client-key,omitempty"`
	ClientKeyPassword *go_netplan_types.NilableString `yaml:"client-key-password,omitempty"`
}
