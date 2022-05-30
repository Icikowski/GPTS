package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

func TestGetLoggerForRouteAndRequest(t *testing.T) {
	routePath, routeType := "/test", "dummy"
	request := &http.Request{
		Method:     http.MethodGet,
		RemoteAddr: "127.0.0.1:8080",
		URL: &url.URL{
			Scheme: "https",
			Host:   "localhost",
			Path:   routePath,
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	baseLogger := zerolog.New(buffer)

	log := getLoggerForRouteAndRequest(baseLogger, routePath, routeType, request)
	log.Info().Msg("test")

	parsed := make(map[string]any, 0)
	json.NewDecoder(buffer).Decode(&parsed)

	for _, key := range []string{"route", "request"} {
		require.Containsf(t, parsed, key, "log does not contain %s", key)
	}

	routeNode := parsed["route"].(map[string]any)
	for _, key := range []string{"path", "type"} {
		require.Containsf(t, routeNode, key, "log for route does not contain %s", key)
	}

	requestNode := parsed["request"].(map[string]any)
	for _, key := range []string{"remote", "method", "path"} {
		require.Containsf(t, requestNode, key, "log for request does not contain %s", key)
	}
}

func TestGetResponseForError(t *testing.T) {
	tests := map[string]struct {
		err            error
		expectedStatus int
	}{
		"generic error": {
			err:            errors.New("some error occurred"),
			expectedStatus: http.StatusBadRequest,
		},
		"unsupported media type": {
			err:            errors.New(common.MsgContentTypeNotAllowed),
			expectedStatus: http.StatusUnsupportedMediaType,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualStatus, _ := getResponseForError(tc.err)
			require.Equal(t, tc.expectedStatus, actualStatus, "statuses does not match")
		})
	}
}

func TestGetSortedRoutes(t *testing.T) {
	tests := map[string]struct {
		givenEntries  map[string]config.Route
		expectedPaths []string
	}{
		"scenario 1": {
			givenEntries: map[string]config.Route{
				"/a": {},
				"/b": {},
				"/c": {},
			},
			expectedPaths: []string{"/c", "/b", "/a"},
		},
		"scenario 2": {
			givenEntries: map[string]config.Route{
				"/":    {},
				"/a":   {},
				"/b":   {},
				"/a/x": {},
			},
			expectedPaths: []string{"/a/x", "/b", "/a", "/"},
		},
		"scenario 3": {
			givenEntries: map[string]config.Route{
				"/a":   {},
				"/":    {},
				"/b":   {},
				"/b/x": {},
				"/a/x": {},
			},
			expectedPaths: []string{"/b/x", "/a/x", "/b", "/a", "/"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualPaths := getSortedRoutes(tc.givenEntries)
			require.EqualValues(t, tc.expectedPaths, actualPaths, "paths sorted incorrectly")
		})
	}
}
