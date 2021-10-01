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

func main() {
	log := logger.ForComponent("cli")

	log.Info().
		Str("serverPort", common.ServerPort).
		Str("healthchecksPort", common.HealthcheckPort).
		Msg("starting application")

	healthServer := health.PrepareHealthEndpoints(common.HealthcheckPort)
	go func() {
		log.Info().Msg("health endpoints starting")
		if err := healthServer.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("health endpoints have been shut down unexpectedly")
		}
	}()

	if common.DefaultConfigOnStartup == "true" {
		config.CurrentConfiguration.SetDefaultConfiguration()
	}

	log.Info().Msg("marking application liveness as UP")
	health.ApplicationStatus.SetStatus(true)

	for {
		service.ExpectingShutdown = false
		server := service.PrepareServer(common.ServerPort)
		health.TestServiceStatus.SetStatus(true)
		if err := server.ListenAndServe(); err != nil {
			if service.ExpectingShutdown && err == http.ErrServerClosed {
				log.Info().Msg("service has been shut down for planned maintenance")
			} else {
				log.Fatal().Err(err).Msg("service has been shut down unexpectedly")
			}
		}
	}
}
