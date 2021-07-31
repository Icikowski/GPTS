package config

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
)

func TestGetConfigurationAndSetConfiguration(t *testing.T) {
	routes := []RouteDefinition{
		{ID: "a", Path: "/a", Method: common.MethodAll},
		{ID: "b_get", Path: "/b", Method: http.MethodGet},
		{ID: "b_post", Path: "/b", Method: http.MethodPost},
		{ID: "b_put", Path: "/b", Method: http.MethodPut},
	}
	config := configuration{}

	require.Empty(t, config.GetConfiguration(), "configuration not properly retrieved")

	config.SetConfiguration(routes)
	require.EqualValues(t, routes, config.routes, "configuration not properly set")

	require.EqualValues(t, routes, config.GetConfiguration(), "configuration not properly retrieved")
}
