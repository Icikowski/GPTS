package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

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
