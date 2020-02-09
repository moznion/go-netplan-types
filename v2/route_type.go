package netplan

// RouteType represents netplan's route type as nillable.
type RouteType struct {
	val        routeType
	isAssigned bool
}

type routeType string

const (
	unicastRouteType     routeType = "unicast"
	unreachableRouteType           = "unreachable"
	blackholeRouteType             = "blackhole"
	prohibitRouteType              = "prohibit"
)

// UnicastRouteType returns `unicast` route type.
func UnicastRouteType() *RouteType {
	return &RouteType{
		val:        unicastRouteType,
		isAssigned: true,
	}
}

// UnreachableRouteType returns `unreachable` route type.
func UnreachableRouteType() *RouteType {
	return &RouteType{
		val:        unreachableRouteType,
		isAssigned: true,
	}
}

// BlackholeRouteType returns `blackhole` route type
func BlackholeRouteType() *RouteType {
	return &RouteType{
		val:        blackholeRouteType,
		isAssigned: true,
	}
}

// ProhibitRouteType returns `prohibit` route type
func ProhibitRouteType() *RouteType {
	return &RouteType{
		val:        prohibitRouteType,
		isAssigned: true,
	}
}

// MarshalYAML marshals RouteType as YAML.
// This method used on marshaling YAML internally.
func (rt *RouteType) MarshalYAML() (interface{}, error) {
	if rt.isAssigned {
		return rt.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals RouteType as YAML.
// This method used on unmarshaling YAML internally.
func (rt *RouteType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val routeType
	if err := unmarshal(&val); err != nil {
		return err
	}

	rt.val = val
	rt.isAssigned = true

	return nil
}
