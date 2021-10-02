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
	logger.InitializeLog()
}

var version = common.BuildValueUnknown
var gitCommit = common.BuildValueUnknown

func main() {
	log := logger.ForComponent("cli")

	log.Info().
		Str("version", version).
		Str("gitCommit", gitCommit).
		Msg("version information")

	log.Info().
		Str("servicePort", common.ServicePort).
		Str("healthchecksPort", common.HealthchecksPort).
		Msg("starting application")

	healthServer := health.PrepareHealthEndpoints(common.HealthchecksPort)
	go func() {
		log.Debug().Msg("health endpoints starting")
		if err := healthServer.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("health endpoints have been shut down unexpectedly")
		}
	}()

	if common.DefaultConfigOnStartup == "true" {
		config.CurrentConfiguration.SetDefaultConfiguration()
	}

	log.Debug().Msg("marking application liveness as UP")
	health.ApplicationStatus.SetStatus(true)

	for {
		service.ExpectingShutdown = false
		server := service.PrepareServer(common.ServicePort)
		health.ServiceStatus.SetStatus(true)
		if err := server.ListenAndServe(); err != nil {
			if service.ExpectingShutdown && err == http.ErrServerClosed {
				log.Info().Msg("service has been shut down for planned maintenance")
			} else {
				log.Fatal().Err(err).Msg("service has been shut down unexpectedly")
			}
		}
	}
}
