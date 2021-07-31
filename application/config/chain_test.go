package config

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/utils"
)

func TestNewRoutesChain(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	routesChain := newRoutesChain(routes)

	require.EqualValues(t, routes, *routesChain.currentState)
}

func TestRoutesChainFilter(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	tests := map[string]struct {
		filterFunc     func(RouteDefinition) bool
		expectedResult []RouteDefinition
	}{
		"all pass": {
			filterFunc: func(_ RouteDefinition) bool {
				return true
			},
			expectedResult: []RouteDefinition{
				{ID: "a", Path: "/a", Method: common.MethodAll},
				{ID: "b_get", Path: "/b", Method: http.MethodGet},
				{ID: "b_post", Path: "/b", Method: http.MethodPost},
				{ID: "b_put", Path: "/b", Method: http.MethodPut},
			},
		},
		"none pass": {
			filterFunc: func(_ RouteDefinition) bool {
				return false
			},
			expectedResult: []RouteDefinition{},
		},
		"non-global pass": {
			filterFunc: func(rd RouteDefinition) bool {
				return rd.Method == common.MethodAll
			},
			expectedResult: []RouteDefinition{
				{ID: "a", Path: "/a", Method: common.MethodAll},
			},
		},
		"chosen path pass": {
			filterFunc: func(rd RouteDefinition) bool {
				return rd.Path == "/b"
			},
			expectedResult: []RouteDefinition{
				{ID: "b_get", Path: "/b", Method: http.MethodGet},
				{ID: "b_post", Path: "/b", Method: http.MethodPost},
				{ID: "b_put", Path: "/b", Method: http.MethodPut},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			routesChain := newRoutesChain(routes).
				filter(tc.filterFunc)
			require.EqualValues(t, tc.expectedResult, *routesChain.currentState)
		})
	}
}

func TestRoutesChainTransform(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	tests := map[string]struct {
		transformFunc  func(RouteDefinition) RouteDefinition
		expectedResult []RouteDefinition
	}{
		"change status to Unauthorized": {
			transformFunc: func(rd RouteDefinition) RouteDefinition {
				rd.Status = utils.IntToPointer(http.StatusUnauthorized)
				return rd
			},
			expectedResult: []RouteDefinition{
				{ID: "a", Path: "/a", Method: common.MethodAll, Status: utils.IntToPointer(http.StatusUnauthorized)},
				{ID: "b_get", Path: "/b", Method: http.MethodGet, Status: utils.IntToPointer(http.StatusUnauthorized)},
				{ID: "b_post", Path: "/b", Method: http.MethodPost, Status: utils.IntToPointer(http.StatusUnauthorized)},
				{ID: "b_put", Path: "/b", Method: http.MethodPut, Status: utils.IntToPointer(http.StatusUnauthorized)},
			},
		},
		"remove IDs": {
			transformFunc: func(rd RouteDefinition) RouteDefinition {
				rd.ID = ""
				return rd
			},
			expectedResult: []RouteDefinition{
				{Path: "/a", Method: common.MethodAll},
				{Path: "/b", Method: http.MethodGet},
				{Path: "/b", Method: http.MethodPost},
				{Path: "/b", Method: http.MethodPut},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			routesChain := newRoutesChain(routes).
				transform(tc.transformFunc)
			require.EqualValues(t, tc.expectedResult, *routesChain.currentState)
		})
	}
}

func TestRoutesChainWithSameBaseSpecAs(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	tests := map[string]struct {
		example        RouteDefinition
		expectedResult []RouteDefinition
	}{
		"match": {
			example: RouteDefinition{Path: "/a", Method: common.MethodAll},
			expectedResult: []RouteDefinition{
				{ID: "a", Path: "/a", Method: common.MethodAll},
			},
		},
		"no match": {
			example:        RouteDefinition{Path: "/a", Method: http.MethodPost},
			expectedResult: []RouteDefinition{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			routesChain := newRoutesChain(routes).
				withSameBaseSpecAs(tc.example)
			require.EqualValues(t, tc.expectedResult, *routesChain.currentState)
		})
	}
}

func TestRoutesChainGetAll(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	require.EqualValues(t, routes, *newRoutesChain(routes).getAll())
}

func TestRoutesChainGetOne(t *testing.T) {
	tests := map[string]struct {
		initialState  []RouteDefinition
		expectedRoute *RouteDefinition
		shouldFail    bool
	}{
		"single route present": {
			initialState: []RouteDefinition{
				{Path: "/a", Method: http.MethodGet},
			},
			expectedRoute: &RouteDefinition{Path: "/a", Method: http.MethodGet},
			shouldFail:    false,
		},
		"multiple routes present": {
			initialState: []RouteDefinition{
				{Path: "/a", Method: http.MethodGet},
				{Path: "/b", Method: http.MethodGet},
			},
			expectedRoute: nil,
			shouldFail:    true,
		},
		"no route present": {
			initialState:  []RouteDefinition{},
			expectedRoute: nil,
			shouldFail:    false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := newRoutesChain(tc.initialState).getOne()

			if tc.shouldFail {
				require.Nil(t, result)
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			if tc.expectedRoute == nil {
				require.Nil(t, result)
				return
			}
			require.EqualValues(t, *tc.expectedRoute, *result)
		})
	}
}

func TestRoutesChainGetPaths(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	require.EqualValues(t,
		[]string{"/a", "/b", "/b", "/b"},
		*newRoutesChain(routes).getPaths(),
	)
}

func TestRoutesChainGetMethods(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}

	require.EqualValues(t,
		[]string{common.MethodAll, http.MethodGet, http.MethodPost, http.MethodPut},
		*newRoutesChain(routes).getMethods(),
	)
}
