package netplan

import (
	"fmt"
	"strconv"
	"strings"

	yamlnillable "github.com/moznion/go-yaml-nillable"
)

// Address represents an IP address.
// If `PrefixLen` existed, YAML marshaled value will be `{Address}/{PrefixLen}`,
// on the other hand, once `PrefixLen` hasn't existed the marshaled value will be `{Address}`.
type Address struct {
	Address   string
	PrefixLen *yamlnillable.Uint8
}

// MarshalYAML marshals Address as YAML.
// This method used on marshaling YAML internally.
func (addr *Address) MarshalYAML() (interface{}, error) {
	if addr.PrefixLen == nil || !addr.PrefixLen.IsAssigned {
		return addr.Address, nil
	}

	return fmt.Sprintf("%s/%d", addr.Address, addr.PrefixLen.Val), nil
}

// UnmarshalYAML unmarshals Address as YAML.
// This method used on unmarshaling YAML internally.
func (addr *Address) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}

	ss := strings.Split(val, "/")
	if len(ss) <= 0 {
		return nil
	}

	addr.Address = ss[0]
	if len(ss) >= 2 {
		prefixLen, err := strconv.ParseUint(ss[1], 10, 8)
		if err != nil {
			return err
		}
		addr.PrefixLen = yamlnillable.Uint8Of(uint8(prefixLen))
	}

	return nil
}
