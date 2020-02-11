package netplan

import (
	"testing"

	yamlnillable "github.com/moznion/go-yaml-nillable"
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
		Password:          yamlnillable.StringOf("pswd"),
		Method:            TLSAuthMethod(),
		Identity:          yamlnillable.StringOf("ident"),
		AnonymousIdentity: yamlnillable.StringOf("anon-ident"),
		CACertificate:     yamlnillable.StringOf("ca-cert"),
		ClientCertificate: yamlnillable.StringOf("client-cert"),
		ClientKey:         yamlnillable.StringOf("client-key"),
		ClientKeyPassword: yamlnillable.StringOf("client-key-pswd"),
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
