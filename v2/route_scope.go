package netplan

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

func GlobalRouteScope() *RouteScope {
	return &RouteScope{
		val:        globalRouteScope,
		isAssigned: true,
	}
}

func LinkRouteScope() *RouteScope {
	return &RouteScope{
		val:        linkRouteScope,
		isAssigned: true,
	}
}

func HostRouteScope() *RouteScope {
	return &RouteScope{
		val:        hostRouteScope,
		isAssigned: true,
	}
}

func (rs *RouteScope) MarshalYAML() (interface{}, error) {
	if rs.isAssigned {
		return rs.val, nil
	}
	return nil, nil
}

func (rs *RouteScope) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val routeScope
	if err := unmarshal(&val); err != nil {
		return err
	}

	rs.val = val
	rs.isAssigned = true

	return nil
}
