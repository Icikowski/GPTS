package service

// TODO: Adapt tests to new architecture

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/config"
	"icikowski.pl/gpts/utils"
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
			expectedKeys:   &[]string{"/hello", "/no-details", "/base64", "/bad-base64", "/multiple-methods"},
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
			expectedStatus: http.StatusServiceUnavailable,
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
			expectedKeys:   &[]string{"host", "path", "method", "headers"},
		},
		"get default endpoint on other path": {
			method:         http.MethodGet,
			path:           "/a/b/c",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers"},
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
	}

	config.CurrentConfiguration.SetConfiguration(map[string]config.Route{})
	config.CurrentConfiguration.SetDefaultConfiguration()

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
			Content:     utils.StringToPointer(`{"get": true}`),
			ContentType: utils.StringToPointer("application/json"),
		},
		POST: &config.Response{
			Content:     utils.StringToPointer(`{"post": true}`),
			ContentType: utils.StringToPointer("application/json"),
		},
		PUT: &config.Response{
			Content:     utils.StringToPointer(`{"put": true}`),
			ContentType: utils.StringToPointer("application/json"),
		},
		PATCH: &config.Response{
			Content:     utils.StringToPointer(`{"patch": true}`),
			ContentType: utils.StringToPointer("application/json"),
		},
		DELETE: &config.Response{
			Content:     utils.StringToPointer(`{"delete": true}`),
			ContentType: utils.StringToPointer("application/json"),
		},
	}

	config.CurrentConfiguration.SetConfiguration(extendedConfig)

	server := PrepareServer("30101")
	go func() {
		_ = server.ListenAndServe()
	}()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{}

			request, err := http.NewRequest(tc.method, "http://localhost:30101"+tc.path, nil)
			require.NoError(t, err, "no error expected during request composition")
			response, err := client.Do(request)
			require.NoError(t, err, "no error expected")

			require.Equal(t, tc.expectedStatus, response.StatusCode, "status code is different than expected")

			if tc.expectedKeys != nil {
				var output map[string]interface{}
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

	err := server.Shutdown(context.Background())
	require.NoError(t, err, "server not closed properly")
}
