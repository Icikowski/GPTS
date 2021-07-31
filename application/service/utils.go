package service

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

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

func getSortedRoutes(routes []config.RouteDefinition) []config.RouteDefinition {
	sorted := make([]config.RouteDefinition, len(routes))
	copy(sorted, routes)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Path > sorted[j].Path
	})
	sort.SliceStable(sorted, func(i, j int) bool {
		return strings.Count(sorted[i].Path, "/") > strings.Count(sorted[j].Path, "/")
	})
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Method > sorted[j].Method
	})
	return sorted
}
