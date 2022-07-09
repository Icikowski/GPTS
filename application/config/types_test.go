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

func TestGetResponseForMethod(t *testing.T) {
	emptyResponse := &Response{}
	fullRoute := Route{
		GET:    emptyResponse,
		PATCH:  emptyResponse,
		POST:   emptyResponse,
		PUT:    emptyResponse,
		DELETE: emptyResponse,
	}
	strippedRoute := Route{
		GET:     emptyResponse,
		Default: emptyResponse,
	}
	singleMethodRoute := Route{
		GET: emptyResponse,
	}

	tests := map[string]struct {
		examinedRoute    Route
		method           string
		expectedResponse *Response
	}{
		"get GET response with match": {
			examinedRoute:    fullRoute,
			method:           http.MethodGet,
			expectedResponse: emptyResponse,
		},
		"get POST response with match": {
			examinedRoute:    fullRoute,
			method:           http.MethodPost,
			expectedResponse: emptyResponse,
		},
		"get PUT response with match": {
			examinedRoute:    fullRoute,
			method:           http.MethodPatch,
			expectedResponse: emptyResponse,
		},
		"get PATCH response with match": {
			examinedRoute:    fullRoute,
			method:           http.MethodPut,
			expectedResponse: emptyResponse,
		},
		"get DELETE response with match": {
			examinedRoute:    fullRoute,
			method:           http.MethodDelete,
			expectedResponse: emptyResponse,
		},
		"get POST response with default fallback": {
			examinedRoute:    strippedRoute,
			method:           http.MethodPost,
			expectedResponse: emptyResponse,
		},
		"get POST response with no default fallback": {
			examinedRoute:    singleMethodRoute,
			method:           http.MethodPost,
			expectedResponse: nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectedResponse, tc.examinedRoute.GetResponseForMethod(tc.method))
		})
	}
}
