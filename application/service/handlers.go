package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
)

func getConfigHandlerFunction(server *http.Server) func(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("preparing configuration handler")
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
					Str("path", "/config").
					Str("type", "builtin"),
			).
			Logger()

		resolveContent := func(target interface{}) error {
			var resolverError error
			switch r.Header.Get("Content-Type") {
			case common.ContentTypeJSON:
				innerLog.Debug().Msg("detected JSON configuration payload")
				resolverError = getDecoder(common.ContentTypeJSON, r.Body)(target)
			case common.ContentTypeYAML:
				innerLog.Debug().Msg("detected YAML configuration payload")
				resolverError = getDecoder(common.ContentTypeYAML, r.Body)(target)
			default:
				innerLog.Debug().Msg("Detected unknown payload type")
				resolverError = errors.New(common.MsgContentTypeNotAllowed)
			}

			if resolverError != nil {
				innerLog.Warn().Err(resolverError).Msg("error while parsing configuration")
				status, payload := getResponseForError(resolverError)
				w.Header().Set("Content-Type", common.ContentTypeJSON)
				w.WriteHeader(status)
				_, _ = w.Write(payload)
				return resolverError
			}
			return nil
		}

		sendAcceptanceAndScheduleServiceShutdown := func() {
			go func() {
				ExpectingShutdown = true
				innerLog.Info().Msg("scheduling service stop")
				time.Sleep(2 * time.Second)
				_ = server.Shutdown(context.Background())
			}()
			w.WriteHeader(http.StatusAccepted)
		}

		var err error

		switch r.Method {
		case http.MethodGet:
			mediaType := r.Header.Get("Accept")
			if strings.Contains(mediaType, "yaml") {
				mediaType = common.ContentTypeYAML
			} else {
				mediaType = common.ContentTypeJSON
			}

			w.Header().Set("Content-Type", mediaType)
			w.WriteHeader(http.StatusOK)
			_ = getEncoder(mediaType, w)(config.CurrentConfiguration.GetConfiguration())

			innerLog.Info().Msg("returned current configuration")
			return
		case http.MethodPost:
			innerLog.Info().Msg("processing incoming configuration")

			var entries []config.RouteDefinition
			err = resolveContent(&entries)
			if err != nil {
				return
			}

			valid, warns := config.CurrentConfiguration.ValidateConfiguration(entries)
			if !valid {
				innerLog.Warn().Strs("warnings", *warns).Msg("configuration cannot be applied due to warnings")
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(*warns)
				return
			}

			config.CurrentConfiguration.SetConfiguration(entries)
			sendAcceptanceAndScheduleServiceShutdown()
			innerLog.Info().Msg("configuration applied")
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			innerLog.Warn().Msg("method not allowed")
		}
	}
}

func getDefaultHandler() http.Handler {
	log.Info().Msg("preparing default handler")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
					Str("path", "*").
					Str("type", "builtin"),
			).
			Logger()

		response := defaultResponse{
			Host:    r.Host,
			Path:    r.URL.Path,
			Method:  r.Method,
			Headers: r.Header,
		}

		mediaType := r.Header.Get("Accept")
		if strings.Contains(mediaType, "yaml") {
			mediaType = common.ContentTypeYAML
		} else {
			mediaType = common.ContentTypeJSON
		}

		w.Header().Add("Content-Type", mediaType)
		w.WriteHeader(http.StatusOK)
		_ = getEncoder(mediaType, w)(response)

		innerLog.Info().Msg("request served")
	})
}
