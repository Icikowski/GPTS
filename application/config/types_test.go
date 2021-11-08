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
		Status:      utils.IntToPointer(http.StatusCreated),
		ContentType: utils.StringToPointer(common.ContentTypeJSON),
		Headers:     &map[string]string{"a": "b"},
	}

	testLog.Info().Object("response", &response).Send()

	var contents map[string]interface{}
	err := json.Unmarshal(buffer.Bytes(), &contents)
	require.NoError(t, err, "no error expected during reading logged Response")

	require.EqualValues(t, map[string]interface{}{
		"level": "info",
		"response": map[string]interface{}{
			"configured":  true,
			"status":      float64(http.StatusCreated),
			"contentType": common.ContentTypeJSON,
			"headers": map[string]interface{}{
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

	var contents map[string]interface{}
	err := json.Unmarshal(buffer.Bytes(), &contents)
	require.NoError(t, err, "no error expected during reading logged Route")

	require.EqualValues(t, map[string]interface{}{
		"level": "info",
		"route": map[string]interface{}{
			"allowSubpaths": true,
			"methods": map[string]interface{}{
				"default": map[string]interface{}{"configured": true},
				"get":     map[string]interface{}{"configured": true},
				"post":    map[string]interface{}{"configured": true},
				"put":     map[string]interface{}{"configured": false},
				"patch":   map[string]interface{}{"configured": false},
				"delete":  map[string]interface{}{"configured": false},
			},
		},
	}, contents, "log contains unexpected Route properties")
}
