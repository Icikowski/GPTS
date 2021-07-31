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
		givenEntries  []config.RouteDefinition
		expectedPaths []config.RouteDefinition
	}{
		"scenario 1": {
			givenEntries: []config.RouteDefinition{
				{Path: "/a"},
				{Path: "/b"},
				{Path: "/c"},
			},
			expectedPaths: []config.RouteDefinition{
				{Path: "/c"},
				{Path: "/b"},
				{Path: "/a"},
			},
		},
		"scenario 2": {
			givenEntries: []config.RouteDefinition{
				{Path: "/"},
				{Path: "/a"},
				{Path: "/b"},
				{Path: "/a/x"},
			},
			expectedPaths: []config.RouteDefinition{
				{Path: "/a/x"},
				{Path: "/b"},
				{Path: "/a"},
				{Path: "/"},
			},
		},
		"scenario 3": {
			givenEntries: []config.RouteDefinition{
				{Path: "/a"},
				{Path: "/"},
				{Path: "/b"},
				{Path: "/b/x"},
				{Path: "/a/x"},
			},
			expectedPaths: []config.RouteDefinition{
				{Path: "/b/x"},
				{Path: "/a/x"},
				{Path: "/b"},
				{Path: "/a"},
				{Path: "/"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualPaths := getSortedRoutes(tc.givenEntries)
			require.EqualValues(t, tc.expectedPaths, actualPaths, "paths sorted incorrectly")
		})
	}
}
