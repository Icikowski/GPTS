package config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/utils"
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

func TestSetDefaultConfiguration(t *testing.T) {
	config := configuration{}

	require.Empty(t, config.GetConfiguration(), "configuration not properly retrieved")

	config.SetDefaultConfiguration(zerolog.Nop())
	require.Contains(t, config.routes, "/hello", "default configuration not properly set")
}

func TestResponseMarshalZerologObject(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	testLog := zerolog.New(buffer)

	response := Response{
		Status:      utils.PointerTo(http.StatusCreated),
		ContentType: utils.PointerTo(common.ContentTypeJSON),
		Headers:     &map[string]string{"a": "b"},
	}

	testLog.Info().Object("response", &response).Send()

	var contents map[string]any
	err := json.Unmarshal(buffer.Bytes(), &contents)
	require.NoError(t, err, "no error expected during reading logged Response")

	require.EqualValues(t, map[string]any{
		"level": "info",
		"response": map[string]any{
			"configured":  true,
			"status":      float64(http.StatusCreated),
			"contentType": common.ContentTypeJSON,
			"headers": map[string]any{
				"a": "b",
			},
		},
	}, contents, "log contains unexpected Response properties")
}

func TestRouteMarshalZerologObject(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	testLog := zerolog.New(buffer)

	route := Route{
		AllowSubpaths: true,
		GET:           &Response{},
		POST:          &Response{},
		Default:       &Response{},
	}

	testLog.Info().Object("route", route).Send()

	var contents map[string]any
	err := json.Unmarshal(buffer.Bytes(), &contents)
	require.NoError(t, err, "no error expected during reading logged Route")

	require.EqualValues(t, map[string]any{
		"level": "info",
		"route": map[string]any{
			"allowSubpaths": true,
			"methods": map[string]any{
				"default": map[string]any{"configured": true},
				"get":     map[string]any{"configured": true},
				"post":    map[string]any{"configured": true},
				"put":     map[string]any{"configured": false},
				"patch":   map[string]any{"configured": false},
				"delete":  map[string]any{"configured": false},
			},
		},
	}, contents, "log contains unexpected Route properties")
}
