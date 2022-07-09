package service

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
	"icikowski.pl/gpts/health"
	"icikowski.pl/gpts/utils"
)

func getHandlerForRoute(path string, route config.Route, log zerolog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		innerLog := log.With().
			Dict(
				"request",
				zerolog.Dict().
					Str("remote", r.RemoteAddr).
					Str("path", r.URL.Path).
					Str("method", r.Method),
			).
			Dict(
				"endpoint",
				zerolog.Dict().
					Str("path", path).
					Str("type", "user"),
			).
			Logger()

		var (
			status      *int
			contentType *string
			content     *string
			headers     *map[string]string
		)

		if response := route.GetResponseForMethod(r.Method); response != nil {
			status = response.Status
			contentType = response.ContentType
			content = response.Content
			headers = response.Headers
		} else {
			status = utils.PointerTo(http.StatusMethodNotAllowed)
		}

		var finalContent []byte

		if headers != nil {
			for headerName, headerValue := range *headers {
				w.Header().Set(headerName, headerValue)
			}
		}

		status = utils.GetOptionalOrFallback(status, utils.PointerTo(http.StatusOK))
		contentType = utils.GetOptionalOrFallback(contentType, utils.PointerTo("text/plain"))
		content = utils.GetOptionalOrFallback(content, new(string))

		if strings.HasPrefix(strings.TrimSpace(*content), "base64,") {
			var err error
			finalContent, err = base64.StdEncoding.DecodeString(strings.TrimPrefix(strings.TrimSpace(*content), "base64,"))
			if err != nil {
				innerLog.Warn().Err(err).Msg("cannot decode base64 content")
				*status = http.StatusInternalServerError
			}
		} else {
			finalContent = []byte(*content)
		}

		w.Header().Set(common.HeaderContentType, *contentType)
		w.WriteHeader(*status)
		_, _ = w.Write(finalContent)
		innerLog.Info().Msg("request served")
	}
}

// PrepareServer prepares, configures and runs test service server
func PrepareServer(log zerolog.Logger, port int) *http.Server {
	log.Info().Msg("preparing test service's router & server")

	r := mux.NewRouter().StrictSlash(true)
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", port),
	}

	r.HandleFunc(common.ConfigurationEndpoint, getConfigHandlerFunction(log, server))

	entries := config.CurrentConfiguration.GetConfiguration()
	sortedRoutes := getSortedRoutes(entries)
	log.Debug().Msg("paths registration order determined")
	for _, path := range sortedRoutes {
		path := path
		route := entries[path]

		log.Info().
			Str("path", path).
			Object("route", route).
			Msg("preparing handler")

		handler := getHandlerForRoute(path, route, log)

		if route.AllowSubpaths {
			r.PathPrefix(path).HandlerFunc(handler)
		} else {
			r.HandleFunc(path, handler)
		}
	}

	r.NotFoundHandler = getDefaultHandler(log)

	log.Debug().Msg("registering shutdown hooks")
	server.RegisterOnShutdown(health.ServiceStatus.MarkAsDown)

	log.Info().Msg("server prepared")
	return server
}
