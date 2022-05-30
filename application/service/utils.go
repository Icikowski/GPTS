package service

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

func getLoggerForRouteAndRequest(log zerolog.Logger, routePath, routeType string, request *http.Request) zerolog.Logger {
	return log.With().
		Dict(
			"route",
			zerolog.Dict().
				Str("path", routePath).
				Str("type", routeType),
		).
		Dict(
			"request",
			zerolog.Dict().
				Str("remote", request.RemoteAddr).
				Str("method", request.Method).
				Str("path", request.URL.Path),
		).
		Logger()
}

func getResponseForError(err error) (int, []byte) {
	var details struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Details string `json:"details,omitempty"`
	}

	if err.Error() == common.MsgContentTypeNotAllowed {
		details.Status = http.StatusUnsupportedMediaType
		details.Message = err.Error()
	} else {
		details.Status = http.StatusBadRequest
		details.Message = "configuration can't be parsed"
		details.Details = err.Error()
	}

	payload, _ := json.Marshal(details)
	return details.Status, payload
}

func getSortedRoutes(routes map[string]config.Route) []string {
	sorted := []string{}
	for key := range routes {
		sorted = append(sorted, key)
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})
	sort.SliceStable(sorted, func(i, j int) bool {
		return strings.Count(sorted[i], "/") > strings.Count(sorted[j], "/")
	})
	return sorted
}
