package netplan

// TunnelMode represents netplan's tunnel mode as nillable.
type TunnelMode struct {
	val        tunnelMode
	isAssigned bool
}

type tunnelMode string

const (
	sitTunnelMode       tunnelMode = "sit"
	greTunnelMode                  = "gre"
	ip6GRETunnelMode               = "ip6gre"
	ipIPTunnelMode                 = "ipip"
	ipIP6TunnelMode                = "ipip6"
	ip6IP6TunnelMode               = "ip6ip6"
	vtiTunnelMode                  = "vti"
	vti6TunnelMode                 = "vti6"
	greTAPTunnelMode               = "gretap"
	ip6GRETAPTunnelMode            = "ip6gretap"
	isatapTunnelMode               = "isatap"
)

// SITTunnelMode returns `sit` tunnel mode.
func SITTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        sitTunnelMode,
		isAssigned: true,
	}
}

// GRETunnelMode returns `gre` tunnel mode.
func GRETunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        greTunnelMode,
		isAssigned: true,
	}
}

// IP6GRETunnelMode returns `ip6gre` tunnel mode.
func IP6GRETunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6GRETunnelMode,
		isAssigned: true,
	}
}

// IPIPTunnelMode returns `ipip` tunnel mode.
func IPIPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ipIPTunnelMode,
		isAssigned: true,
	}
}

// IPIP6TunnelMode returns `ipip6` tunnel mode.
func IPIP6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ipIP6TunnelMode,
		isAssigned: true,
	}
}

// IP6IP6TunnelMode returns `ip6ip6` tunnel mode.
func IP6IP6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6IP6TunnelMode,
		isAssigned: true,
	}
}

// VTITunnelMode returns `vti` tunnel mode.
func VTITunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        vtiTunnelMode,
		isAssigned: true,
	}
}

// VTI6TunnelMode returns `vti6` tunnel mode.
func VTI6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        vti6TunnelMode,
		isAssigned: true,
	}
}

// GRETAPTunnelMode returns `gretap` tunnel mode.
func GRETAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        greTAPTunnelMode,
		isAssigned: true,
	}
}

// IP6GRETAPTunnelMode returns `ip6gretap` tunnel mode.
func IP6GRETAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6GRETAPTunnelMode,
		isAssigned: true,
	}
}

// ISATAPTunnelMode returns `isatap` tunnel mode.
func ISATAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        isatapTunnelMode,
		isAssigned: true,
	}
}

// MarshalYAML marshals TunnelMode as YAML.
// This method used on marshaling YAML internally.
func (tm *TunnelMode) MarshalYAML() (interface{}, error) {
	if tm.isAssigned {
		return tm.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals TunnelMode as YAML.
// This method used on unmarshaling YAML internally.
func (tm *TunnelMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val tunnelMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	tm.val = val
	tm.isAssigned = true

	return nil
}
