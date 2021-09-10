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
