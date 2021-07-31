package config

import "errors"

type routesChain struct {
	currentState *[]RouteDefinition
}

func newRoutesChain(in []RouteDefinition) *routesChain {
	routes := make([]RouteDefinition, len(in))
	copy(routes, in)

	return &routesChain{currentState: &routes}
}

func (rc *routesChain) filter(filterFunc func(RouteDefinition) bool) *routesChain {
	filtered := []RouteDefinition{}
	for _, route := range *rc.currentState {
		if filterFunc(route) {
			filtered = append(filtered, route)
		}
	}
	rc.currentState = &filtered

	return rc
}

func (rc *routesChain) transform(transformFunc func(RouteDefinition) RouteDefinition) *routesChain {
	transformed := []RouteDefinition{}
	for _, route := range *rc.currentState {
		transformed = append(transformed, transformFunc(route))
	}
	rc.currentState = &transformed

	return rc
}

func (rc *routesChain) withSameBaseSpecAs(example RouteDefinition) *routesChain {
	return rc.filter(func(route RouteDefinition) bool {
		return route.Path == example.Path && route.Method == example.Method
	})
}

func (rc *routesChain) getAll() *[]RouteDefinition {
	return rc.currentState
}

func (rc *routesChain) getOne() (*RouteDefinition, error) {
	if len(*rc.currentState) == 0 {
		return nil, nil
	}

	if len(*rc.currentState) == 1 {
		return &(*rc.currentState)[0], nil
	}

	return nil, errors.New("found more than one route")
}

func (rc *routesChain) getPaths() *[]string {
	out := []string{}

	for _, route := range *rc.currentState {
		out = append(out, route.Path)
	}

	return &out
}

func (rc *routesChain) getMethods() *[]string {
	out := []string{}

	for _, route := range *rc.currentState {
		out = append(out, route.Method)
	}

	return &out
}
