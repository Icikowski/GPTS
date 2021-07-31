package service

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
	"icikowski.pl/gpts/health"
)

// PrepareServer prepares, configures and runs test service server
func PrepareServer(port string) *http.Server {
	log.Info().Msg("preparing test service's router & server")

	r := mux.NewRouter().StrictSlash(true)
	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	r.HandleFunc("/config", getConfigHandlerFunction(server))

	entries := config.CurrentConfiguration.GetConfiguration()
	sortedRoutes := getSortedRoutes(entries)
	log.Info().Msg("paths registration order determined")
	for _, route := range sortedRoutes {
		log.Info().
			Dict(
				"subject",
				zerolog.Dict().
					Str("method", route.Method).
					Str("path", route.Path),
			).
			Msg("preparing handler")

		var handler = func(w http.ResponseWriter, r *http.Request) {
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
						Str("path", route.Path).
						Str("method", route.Method).
						Str("type", "user"),
				).
				Logger()

			status := route.Status
			contentType := route.ContentType
			content := route.Content
			headers := route.Headers

			var finalContent []byte

			if headers != nil {
				for headerName, headerValue := range *headers {
					w.Header().Set(headerName, headerValue)
				}
			}
			if status == nil {
				status = new(int)
				*status = 200
			}
			if contentType == nil {
				contentType = new(string)
				*contentType = "text/plain"
			}
			if content == nil {
				content = new(string)
			}
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

			w.Header().Set("Content-Type", *contentType)
			w.WriteHeader(*status)
			_, _ = w.Write(finalContent)
			innerLog.Info().Msg("request served")
		}

		rt := r.Get(route.ID)
		if rt == nil {
			rt = r.NewRoute().Name(route.ID).Path(route.Path)
		}

		rt.HandlerFunc(handler)
		if route.Method != common.MethodAll {
			rt.Methods(route.Method)
		}
	}

	r.NotFoundHandler = getDefaultHandler()

	log.Info().Msg("registering shutdown hooks")
	server.RegisterOnShutdown(func() {
		health.TestServiceStatus.SetStatus(false)
	})

	log.Info().Msg("server prepared")
	return server
}
