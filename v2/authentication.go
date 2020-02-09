package netplan

import go_netplan_types "github.com/moznion/go-netplan-types"

type Authentication struct {
	KeyManagement     *KeyManagement                   `yaml:"key-management,omitempty"`
	Password          *go_netplan_types.NillableString `yaml:"password,omitempty"`
	Method            *AuthMethod                      `yaml:"method,omitempty"`
	Identity          *go_netplan_types.NillableString `yaml:"identity,omitempty"`
	AnonymousIdentity *go_netplan_types.NillableString `yaml:"anonymous-identity,omitempty"`
	CACertificate     *go_netplan_types.NillableString `yaml:"ca-certificate,omitempty"`
	ClientCertificate *go_netplan_types.NillableString `yaml:"client-certificate,omitempty"`
	ClientKey         *go_netplan_types.NillableString `yaml:"client-key,omitempty"`
	ClientKeyPassword *go_netplan_types.NillableString `yaml:"client-key-password,omitempty"`
}
