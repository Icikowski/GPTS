package main

import (
	"net/http"

	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
	"icikowski.pl/gpts/health"
	"icikowski.pl/gpts/logger"
	"icikowski.pl/gpts/service"
)

func init() {
	logger.InitializeLog(common.PrettyLog, common.LogLevel)
}

var version = common.BuildValueUnknown
var gitCommit = common.BuildValueUnknown
var binaryType = common.BuildValueUnknown

func main() {
	log := logger.GetLogger()
	l := log.With().Str(common.ComponentField, common.ComponentCLI).Logger()

	l.Info().
		Str("version", version).
		Str("gitCommit", gitCommit).
		Str("binaryType", binaryType).
		Msg("version information")

	l.Info().
		Str("servicePort", common.ServicePort).
		Str("healthchecksPort", common.HealthchecksPort).
		Msg("starting application")

	healthServer := health.PrepareHealthEndpoints(log, common.HealthchecksPort)
	go func() {
		l.Debug().Msg("health endpoints starting")
		if err := healthServer.ListenAndServe(); err != nil {
			l.Fatal().Err(err).Msg("health endpoints have been shut down unexpectedly")
		}
	}()

	if common.DefaultConfigOnStartup == "true" {
		config.CurrentConfiguration.SetDefaultConfiguration(log)
	}

	l.Debug().Msg("marking application liveness as UP")
	health.ApplicationStatus.SetStatus(true)

	for {
		service.ExpectingShutdown = false
		server := service.PrepareServer(log, common.ServicePort)
		health.ServiceStatus.SetStatus(true)
		if err := server.ListenAndServe(); err != nil {
			if service.ExpectingShutdown && err == http.ErrServerClosed {
				l.Info().Msg("service has been shut down for planned maintenance")
			} else {
				l.Fatal().Err(err).Msg("service has been shut down unexpectedly")
			}
		}
	}
}
