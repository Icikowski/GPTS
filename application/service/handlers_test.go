package service

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

const testPort = "30100"

func TestGetConfigHandlerFunction(t *testing.T) {
	tests := map[string]struct {
		requestMethod                string
		contentType                  string
		payload                      string
		expectedStatus               int
		expectedContentType          string
		expectedConfigurationEntries []string
		serverShouldBeClosed         bool
	}{
		"invalid method": {
			requestMethod:        http.MethodPut,
			payload:              "",
			expectedStatus:       http.StatusMethodNotAllowed,
			serverShouldBeClosed: false,
		},
		"get configuration as JSON": {
			requestMethod:        http.MethodGet,
			contentType:          common.ContentTypeJSON,
			expectedStatus:       http.StatusOK,
			expectedContentType:  common.ContentTypeJSON,
			serverShouldBeClosed: false,
		},
		"get configuration as YAML": {
			requestMethod:        http.MethodGet,
			contentType:          common.ContentTypeYAML,
			expectedStatus:       http.StatusOK,
			expectedContentType:  common.ContentTypeYAML,
			serverShouldBeClosed: false,
		},
		"get configuration as other media type": {
			requestMethod:        http.MethodGet,
			contentType:          "text/plain",
			expectedStatus:       http.StatusOK,
			expectedContentType:  common.ContentTypeJSON,
			serverShouldBeClosed: false,
		},
		"post valid configuration as JSON": {
			requestMethod:                http.MethodPost,
			contentType:                  common.ContentTypeJSON,
			payload:                      `[{"path":"/a","method":"all"},{"path":"/b","method":"all"}]`,
			expectedStatus:               http.StatusAccepted,
			expectedConfigurationEntries: []string{"/a", "/b"},
			serverShouldBeClosed:         true,
		},
		"post invalid configuration as JSON": {
			requestMethod:        http.MethodPost,
			contentType:          common.ContentTypeJSON,
			payload:              `["/a","/b"]`,
			expectedStatus:       http.StatusBadRequest,
			expectedContentType:  common.ContentTypeJSON,
			serverShouldBeClosed: false,
		},
		"post valid configuration as YAML": {
			requestMethod:                http.MethodPost,
			contentType:                  common.ContentTypeYAML,
			payload:                      "- path: /a\n  method: all\n- path: /b\n  method: all",
			expectedStatus:               http.StatusAccepted,
			expectedConfigurationEntries: []string{"/a", "/b"},
			serverShouldBeClosed:         true,
		},
		"post invalid configuration as YAML": {
			requestMethod:        http.MethodPost,
			contentType:          common.ContentTypeYAML,
			payload:              "/a: hello\n/b: world",
			expectedStatus:       http.StatusBadRequest,
			expectedContentType:  common.ContentTypeJSON,
			serverShouldBeClosed: false,
		},
		"post configuration as other media type": {
			requestMethod:        http.MethodPost,
			contentType:          "text/plain",
			payload:              "Hello :)",
			expectedStatus:       http.StatusUnsupportedMediaType,
			expectedContentType:  common.ContentTypeJSON,
			serverShouldBeClosed: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			initialEntries := []config.RouteDefinition{}
			config.CurrentConfiguration.SetConfiguration(initialEntries)

			server := &http.Server{
				Addr: ":" + testPort,
			}
			mux := http.NewServeMux()

			handlerFunction := getConfigHandlerFunction(server)

			mux.HandleFunc("/config", handlerFunction)
			server.Handler = mux

			serverClosedSync := sync.Mutex{}
			serverClosed := false
			server.RegisterOnShutdown(func() {
				serverClosedSync.Lock()
				serverClosed = true
				serverClosedSync.Unlock()
			})

			go func() {
				_ = server.ListenAndServe()
			}()

			client := &http.Client{}
			bodyReader := bytes.NewBufferString(tc.payload)

			request, err := http.NewRequest(tc.requestMethod, "http://localhost:"+testPort+"/config", bodyReader)
			require.NoError(t, err, "request could not be prepared")
			if tc.contentType != "" {
				request.Header.Add("Accept", tc.contentType)
				request.Header.Add("Content-Type", tc.contentType)
			}

			response, err := client.Do(request)
			require.NoError(t, err, "no error expected")

			require.Equal(t, tc.expectedStatus, response.StatusCode, "status code does not match expected value")
			if tc.expectedContentType != "" {
				require.Equal(t, tc.expectedContentType, response.Header.Get("Content-Type"), "content type does not match expected value")
			}

			currentConfig := config.CurrentConfiguration.GetConfiguration()
			entries := make([]string, 0, len(currentConfig))
			for _, route := range currentConfig {
				entries = append(entries, route.Path)
			}
			require.ElementsMatch(t, entries, tc.expectedConfigurationEntries, "expected configuration not found")

			if tc.serverShouldBeClosed {
				require.Eventually(t, func() bool {
					serverClosedSync.Lock()
					status := serverClosed
					serverClosedSync.Unlock()
					return status
				}, 5*time.Second, 1*time.Second, "server should be eventually closed")
			} else {
				err = server.Shutdown(context.Background())
				require.NoError(t, err, "server not closed properly")
			}
		})
	}
}

func TestGetDefaultHandler(t *testing.T) {
	tests := map[string]struct {
		contentType         string
		expectedContentType string
	}{
		"no preferred content type specified": {
			expectedContentType: common.ContentTypeJSON,
		},
		"get as JSON": {
			contentType:         common.ContentTypeJSON,
			expectedContentType: common.ContentTypeJSON,
		},
		"get as YAML": {
			contentType:         common.ContentTypeYAML,
			expectedContentType: common.ContentTypeYAML,
		},
		"get as other media type": {
			contentType:         "text/plain",
			expectedContentType: common.ContentTypeJSON,
		},
	}

	server := &http.Server{
		Addr:    ":" + testPort,
		Handler: getDefaultHandler(),
	}
	go func() {
		_ = server.ListenAndServe()
	}()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{}

			request, err := http.NewRequest(http.MethodGet, "http://localhost:"+testPort, nil)
			require.NoError(t, err, "request could not be prepared")
			if tc.contentType != "" {
				request.Header.Add("Accept", tc.contentType)
			}

			response, err := client.Do(request)
			require.NoError(t, err, "no error expected")

			require.Equal(t, http.StatusOK, response.StatusCode, "status code is different than 200 OK")
			require.Equal(t, tc.expectedContentType, response.Header.Get("Content-Type"), "content type does not match expected value")
		})
	}

	err := server.Shutdown(context.Background())
	require.NoError(t, err, "server not closed properly")
}
