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
func PrepareServer(log zerolog.Logger, port string) *http.Server {
	l := log.With().Str(common.ComponentField, common.ComponentService).Logger()
	l.Info().Msg("preparing test service's router & server")

	r := mux.NewRouter().StrictSlash(true)
	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	r.HandleFunc("/config", getConfigHandlerFunction(l, server))

	entries := config.CurrentConfiguration.GetConfiguration()
	sortedRoutes := getSortedRoutes(entries)
	l.Debug().Msg("paths registration order determined")
	for _, path := range sortedRoutes {
		path := path
		route := entries[path]

		l.Info().
			Str("path", path).
			Object("route", route).
			Msg("preparing handler")

		var handler = func(w http.ResponseWriter, r *http.Request) {
			innerLog := l.With().
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

			switch {
			case r.Method == http.MethodGet && route.GET != nil:
				status = route.GET.Status
				contentType = route.GET.ContentType
				content = route.GET.Content
				headers = route.GET.Headers
			case r.Method == http.MethodPost && route.POST != nil:
				status = route.POST.Status
				contentType = route.POST.ContentType
				content = route.POST.Content
				headers = route.POST.Headers
			case r.Method == http.MethodPut && route.PUT != nil:
				status = route.PUT.Status
				contentType = route.PUT.ContentType
				content = route.PUT.Content
				headers = route.PUT.Headers
			case r.Method == http.MethodPatch && route.PATCH != nil:
				status = route.PATCH.Status
				contentType = route.PATCH.ContentType
				content = route.PATCH.Content
				headers = route.PATCH.Headers
			case r.Method == http.MethodDelete && route.DELETE != nil:
				status = route.DELETE.Status
				contentType = route.DELETE.ContentType
				content = route.DELETE.Content
				headers = route.DELETE.Headers
			case route.Default != nil:
				status = route.Default.Status
				contentType = route.Default.ContentType
				content = route.Default.Content
				headers = route.Default.Headers
			default:
				status = new(int)
				*status = http.StatusServiceUnavailable
			}

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

			w.Header().Set(common.HeaderContentType, *contentType)
			w.WriteHeader(*status)
			_, _ = w.Write(finalContent)
			innerLog.Info().Msg("request served")
		}

		if route.AllowSubpaths {
			r.PathPrefix(path).HandlerFunc(handler)
		} else {
			r.HandleFunc(path, handler)
		}
	}

	r.NotFoundHandler = getDefaultHandler(l)

	l.Debug().Msg("registering shutdown hooks")
	server.RegisterOnShutdown(func() {
		health.ServiceStatus.SetStatus(false)
	})

	l.Info().Msg("server prepared")
	return server
}
