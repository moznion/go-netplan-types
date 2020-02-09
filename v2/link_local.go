package netplan

// LinkLocal represents netplan's link-local.
type LinkLocal string

const (
	// IPv4LinkLocal is a link-local value of `ipv4`
	IPv4LinkLocal = "ipv4"
	// IPv6LinkLocal is a link-local value of `ipv6`
	IPv6LinkLocal = "ipv6"
)
