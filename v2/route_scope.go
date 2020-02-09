package netplan

// RouteScope represents netplan's route scope as nillable.
type RouteScope struct {
	val        routeScope
	isAssigned bool
}

type routeScope string

const (
	globalRouteScope routeScope = "global"
	linkRouteScope              = "link"
	hostRouteScope              = "host"
)

// GlobalRouteScope returns `global` route scope.
func GlobalRouteScope() *RouteScope {
	return &RouteScope{
		val:        globalRouteScope,
		isAssigned: true,
	}
}

// LinkRouteScope returns `link` route scope.
func LinkRouteScope() *RouteScope {
	return &RouteScope{
		val:        linkRouteScope,
		isAssigned: true,
	}
}

// HostRouteScope returns `host` route scope.
func HostRouteScope() *RouteScope {
	return &RouteScope{
		val:        hostRouteScope,
		isAssigned: true,
	}
}

// MarshalYAML marshals RouteScope as YAML.
// This method used on marshaling YAML internally.
func (rs *RouteScope) MarshalYAML() (interface{}, error) {
	if rs.isAssigned {
		return rs.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals RouteScope as YAML.
// This method used on unmarshaling YAML internally.
func (rs *RouteScope) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val routeScope
	if err := unmarshal(&val); err != nil {
		return err
	}

	rs.val = val
	rs.isAssigned = true

	return nil
}
