package service

// TODO: Adapt tests to new architecture

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/config"
)

func TestPrepareServer(t *testing.T) {
	tests := map[string]struct {
		path           string
		expectedStatus int
		expectedKeys   *[]string
	}{
		"get configuration": {
			path:           "/config",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"/hello", "/no-details", "/base64", "/bad-base64"},
		},
		"get hello route": {
			path:           "/hello",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"message"},
		},
		"get route with no details": {
			path:           "/no-details",
			expectedStatus: http.StatusServiceUnavailable,
		},
		"get route with base64-encoded body": {
			path:           "/base64",
			expectedStatus: http.StatusOK,
		},
		"get route with incorrect base64-encoded body": {
			path:           "/bad-base64",
			expectedStatus: http.StatusInternalServerError,
		},
		"get default endpoint": {
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers"},
		},
		"get default endpoint on other path": {
			path:           "/a/b/c",
			expectedStatus: http.StatusOK,
			expectedKeys:   &[]string{"host", "path", "method", "headers"},
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

	config.CurrentConfiguration.SetConfiguration(extendedConfig)

	server := PrepareServer("30101")
	go func() {
		_ = server.ListenAndServe()
	}()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{}

			response, err := client.Get("http://localhost:30101" + tc.path)
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
