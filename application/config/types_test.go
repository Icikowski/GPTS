package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfigurationAndSetConfiguration(t *testing.T) {
	routes := map[string]Route{
		"/a": {},
		"/b": {
			GET:  &Response{},
			POST: &Response{},
			PUT:  &Response{},
		},
	}
	config := configuration{}

	require.Empty(t, config.GetConfiguration(), "configuration not properly retrieved")

	config.SetConfiguration(routes)
	require.EqualValues(t, routes, config.routes, "configuration not properly set")

	require.EqualValues(t, routes, config.GetConfiguration(), "configuration not properly retrieved")
}
