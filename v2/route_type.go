package netplan

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

func UnicastRouteType() *RouteType {
	return &RouteType{
		val:        unicastRouteType,
		isAssigned: true,
	}
}

func UnreachableRouteType() *RouteType {
	return &RouteType{
		val:        unreachableRouteType,
		isAssigned: true,
	}
}

func BlackholeRouteType() *RouteType {
	return &RouteType{
		val:        blackholeRouteType,
		isAssigned: true,
	}
}

func ProhibitRouteType() *RouteType {
	return &RouteType{
		val:        prohibitRouteType,
		isAssigned: true,
	}
}

func (rt *RouteType) MarshalYAML() (interface{}, error) {
	if rt.isAssigned {
		return rt.val, nil
	}
	return nil, nil
}

func (rt *RouteType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val routeType
	if err := unmarshal(&val); err != nil {
		return err
	}

	rt.val = val
	rt.isAssigned = true

	return nil
}
