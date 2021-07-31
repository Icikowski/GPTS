package config

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"icikowski.pl/gpts/common"
)

var acceptedMethods []string = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	common.MethodAll,
}

func merge(current []RouteDefinition, incoming []RouteDefinition) (*[]RouteDefinition, *[]string) {
	warnings := []string{}

	// Check for empty entries
	nonEmpty := newRoutesChain(incoming).
		filter(func(route RouteDefinition) bool {
			ok := route.Path != "" && route.Method != ""

			if !ok {
				warnings = append(warnings, "some route is missing path and/or method")
			}
			return ok
		}).getAll()

	// Check for duplicates
	nonDuplicated := []RouteDefinition{}
	for _, route := range *nonEmpty {
		originalRoute, err := newRoutesChain(incoming).
			withSameBaseSpecAs(route).
			getOne()

		if err != nil {
			warnings = append(warnings, fmt.Sprintf("route %s duplicated in incoming configuration", route))
			continue
		}
		nonDuplicated = append(nonDuplicated, *originalRoute)
	}

	// Check for Method correctness
	correctMethod := newRoutesChain(nonDuplicated).
		transform(func(route RouteDefinition) RouteDefinition {
			out := route
			out.Method = strings.ToUpper(route.Method)
			return out
		}).
		filter(func(route RouteDefinition) bool {
			ok := false
			for _, method := range acceptedMethods {
				if route.Method == method {
					ok = true
					break
				}
			}

			if !ok {
				warnings = append(warnings, fmt.Sprintf("route %s has unknown HTTP method", route))
			}
			return ok
		}).getAll()

	// Check for global Method effectiveness agains current configuration
	effectiveAgainstCurrent := newRoutesChain(*correctMethod).
		filter(func(route RouteDefinition) bool {
			if route.Method != common.MethodAll {
				return true
			}

			existingWithSamePath := newRoutesChain(current).
				filter(func(subroute RouteDefinition) bool {
					found, _ := newRoutesChain(*correctMethod).withSameBaseSpecAs(subroute).getOne()
					return found != nil
				}).
				filter(func(subroute RouteDefinition) bool {
					return subroute.Method != common.MethodAll && subroute.Path == route.Path
				}).getMethods()

			if len(*existingWithSamePath) != 0 {
				warnings = append(warnings, fmt.Sprintf(
					"route %s will overlap with existing route definitions with same path and methods %s",
					route,
					strings.Join(*existingWithSamePath, ", "),
				))
				return false
			}
			return true
		}).getAll()

	// Check for non-global Method effectiveness
	global := newRoutesChain(*effectiveAgainstCurrent).
		filter(func(route RouteDefinition) bool {
			return route.Method == common.MethodAll
		}).getPaths()
	effective := newRoutesChain(*correctMethod).
		filter(func(route RouteDefinition) bool {
			if route.Method == common.MethodAll {
				return true
			}

			ok := true
			for _, path := range *global {
				if path == route.Path {
					ok = false
					break
				}
			}

			if !ok {
				warnings = append(warnings, fmt.Sprintf(
					"route %s will be ineffective as route with same path and method set as [%s] is present in configuration",
					route,
					common.MethodAll,
				))
			}
			return ok
		}).getAll()

	// Check for URI correctness
	correctUri := newRoutesChain(*effective).
		filter(func(route RouteDefinition) bool {
			uri, err := url.Parse(route.Path)

			if err != nil {
				warnings = append(warnings, fmt.Sprintf("route %s has invalid path", route))
				return false
			}
			if uri.IsAbs() {
				warnings = append(warnings, fmt.Sprintf("route %s has relative path", route))
				return false
			}
			return true
		}).getAll()

	// Check for existing definitions
	finalRoutes := []RouteDefinition{}
	for _, route := range *correctUri {
		if len(route.ID) == 0 {
			existingRoute, err := newRoutesChain(current).
				withSameBaseSpecAs(route).
				getOne()

			if err != nil {
				warnings = append(warnings, fmt.Sprintf("route %s has multiple instances in current configuration", route))
				continue
			}

			if existingRoute != nil {
				route.ID = existingRoute.ID
			} else {
				route.ID = uuid.NewString()
			}
		}

		finalRoutes = append(finalRoutes, route)
	}

	if len(warnings) != 0 {
		return nil, &warnings
	}
	return &finalRoutes, nil
}
