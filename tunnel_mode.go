package netplan

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

func SITTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        sitTunnelMode,
		isAssigned: true,
	}
}

func GRETunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        greTunnelMode,
		isAssigned: true,
	}
}

func IP6GRETunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6GRETunnelMode,
		isAssigned: true,
	}
}

func IPIPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ipIPTunnelMode,
		isAssigned: true,
	}
}

func IPIP6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ipIP6TunnelMode,
		isAssigned: true,
	}
}

func IP6IP6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6IP6TunnelMode,
		isAssigned: true,
	}
}

func VTITunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        vtiTunnelMode,
		isAssigned: true,
	}
}

func VTI6TunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        vti6TunnelMode,
		isAssigned: true,
	}
}

func GRETAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        greTAPTunnelMode,
		isAssigned: true,
	}
}

func IP6GRETAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        ip6GRETAPTunnelMode,
		isAssigned: true,
	}
}

func ISATAPTunnelMode() *TunnelMode {
	return &TunnelMode{
		val:        isatapTunnelMode,
		isAssigned: true,
	}
}

func (tm *TunnelMode) MarshalYAML() (interface{}, error) {
	if tm.isAssigned {
		return tm.val, nil
	}
	return nil, nil
}

func (tm *TunnelMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val tunnelMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	tm.val = val
	tm.isAssigned = true

	return nil
}
