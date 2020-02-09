package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyAuthentication(t *testing.T) {
	given := Authentication{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Authentication
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeAuthentication(t *testing.T) {
	given := Authentication{
		KeyManagement:     EAPKeyManagement(),
		Password:          go_netplan_types.NilableStringOf("pswd"),
		Method:            TLSAuthMethod(),
		Identity:          go_netplan_types.NilableStringOf("ident"),
		AnonymousIdentity: go_netplan_types.NilableStringOf("anon-ident"),
		CACertificate:     go_netplan_types.NilableStringOf("ca-cert"),
		ClientCertificate: go_netplan_types.NilableStringOf("client-cert"),
		ClientKey:         go_netplan_types.NilableStringOf("client-key"),
		ClientKeyPassword: go_netplan_types.NilableStringOf("client-key-pswd"),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`key-management: eap
password: pswd
method: tls
identity: ident
anonymous-identity: anon-ident
ca-certificate: ca-cert
client-certificate: client-cert
client-key: client-key
client-key-password: client-key-pswd
`), marshal)

	var unmarshal Authentication
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
