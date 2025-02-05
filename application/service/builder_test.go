package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.sr.ht/~icikowski/gpts/config"
	"git.sr.ht/~icikowski/gpts/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestPrepareServer(t *testing.T) {
	tests := map[string]struct {
		method         string
		path           string
		expectedStatus int
		expectedKeys   *[]string
	}{
		"get configuration": {
			method:         http.MethodGet,
			path:           "/config",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"/hello", "/no-details", "/base64", "/bad-base64", "/multiple-methods", "/sub"},
		},
		"get hello route": {
			method:         http.MethodGet,
			path:           "/hello",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"message"},
		},
		"get route with no details": {
			method:         http.MethodGet,
			path:           "/no-details",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		"get route with base64-encoded body": {
			method:         http.MethodGet,
			path:           "/base64",
			expectedStatus: http.StatusOK,
		},
		"get route with incorrect base64-encoded body": {
			method:         http.MethodGet,
			path:           "/bad-base64",
			expectedStatus: http.StatusInternalServerError,
		},
		"get default endpoint": {
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers", "queries"},
		},
		"get default endpoint with some queries": {
			method:         http.MethodGet,
			path:           "/?hello=world",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers", "queries"},
		},
		"get default endpoint on other path": {
			method:         http.MethodGet,
			path:           "/a/b/c",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers", "queries"},
		},
		"get on multiple-methods route": {
			method:         http.MethodGet,
			path:           "/multiple-methods",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"get"},
		},
		"post on multiple-methods route": {
			method:         http.MethodPost,
			path:           "/multiple-methods",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"post"},
		},
		"put on multiple-methods route": {
			method:         http.MethodPut,
			path:           "/multiple-methods",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"put"},
		},
		"patch on multiple-methods route": {
			method:         http.MethodPatch,
			path:           "/multiple-methods",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"patch"},
		},
		"delete on multiple-methods route": {
			method:         http.MethodDelete,
			path:           "/multiple-methods",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"delete"},
		},
		"get on subpath of subpaths-enabled route": {
			method:         http.MethodGet,
			path:           "/sub/x/y/z",
			expectedStatus: http.StatusTeapot,
		},
	}

	config.CurrentConfiguration.SetConfiguration(map[string]config.Route{})
	config.CurrentConfiguration.SetDefaultConfiguration(zerolog.Nop())

	extendedConfig := config.CurrentConfiguration.GetConfiguration()

	extendedConfig["/no-details"] = config.Route{}

	base64content := "base64,SGVsbG8sIFdvcmxkIQ=="
	extendedConfig["/base64"] = config.Route{
		Default: &config.Response{
			Content: &base64content,
		},
	}

	badBase64content := "base64,ThisIsNotValidAtAll123!%"
	extendedConfig["/bad-base64"] = config.Route{
		Default: &config.Response{
			Content: &badBase64content,
		},
	}

	extendedConfig["/multiple-methods"] = config.Route{
		GET: &config.Response{
			Content:     utils.PointerTo(`{"get": true}`),
			ContentType: utils.PointerTo("application/json"),
		},
		POST: &config.Response{
			Content:     utils.PointerTo(`{"post": true}`),
			ContentType: utils.PointerTo("application/json"),
		},
		PUT: &config.Response{
			Content:     utils.PointerTo(`{"put": true}`),
			ContentType: utils.PointerTo("application/json"),
		},
		PATCH: &config.Response{
			Content:     utils.PointerTo(`{"patch": true}`),
			ContentType: utils.PointerTo("application/json"),
		},
		DELETE: &config.Response{
			Content:     utils.PointerTo(`{"delete": true}`),
			ContentType: utils.PointerTo("application/json"),
		},
	}

	extendedConfig["/sub"] = config.Route{
		AllowSubpaths: true,
		Default: &config.Response{
			Status:  utils.PointerTo(http.StatusTeapot),
			Content: utils.PointerTo("I'm a teapot and I'm everywhere"),
		},
	}

	config.CurrentConfiguration.SetConfiguration(extendedConfig)

	server := PrepareServer(zerolog.Nop(), 80)
	testServer := httptest.NewServer(server.Handler)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := testServer.Client()

			request, err := http.NewRequest(tc.method, testServer.URL+tc.path, nil)
			require.NoError(t, err, "no error expected during request composition")
			response, err := client.Do(request)
			require.NoError(t, err, "no error expected")

			require.Equal(t, tc.expectedStatus, response.StatusCode, "status code is different than expected")

			if tc.expectedKeys != nil {
				var output map[string]any
				err = json.NewDecoder(response.Body).Decode(&output)
				require.NoError(t, err, "response should be decoded properly")

				keys := make([]string, 0, len(output))
				for key := range output {
					keys = append(keys, key)
				}
				require.ElementsMatch(t, keys, *tc.expectedKeys, "response has different structure than expected")
			}
		})
	}

	testServer.Close()
}
