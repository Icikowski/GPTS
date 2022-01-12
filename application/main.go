package main

import (
	"flag"
	"net/http"

	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/config"
	"icikowski.pl/gpts/health"
	"icikowski.pl/gpts/logger"
	"icikowski.pl/gpts/service"
)

func init() {
	flag.IntVar(&common.ServicePort, "service-port", common.ServicePort, "Port on which the service will be running")
	flag.IntVar(&common.HealthchecksPort, "health-port", common.HealthchecksPort, "Port on which the healthchecks will be running")
	flag.StringVar(&common.ConfigurationEndpoint, "config-endpoint", common.ConfigurationEndpoint, "Path of the configuration endpoint")
	flag.BoolVar(&common.DefaultConfigOnStartup, "default-config", common.DefaultConfigOnStartup, "Enables loading the default configuration on startup")
	flag.BoolVar(&common.PrettyLog, "pretty-log", common.PrettyLog, "Enables the pretty logger")
	flag.StringVar(&common.LogLevel, "log-level", common.LogLevel, "Global log level; one of [debug, info, warn, error, fatal, panic, trace]")
	flag.Parse()

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
		Int("servicePort", common.ServicePort).
		Int("healthchecksPort", common.HealthchecksPort).
		Str("configurationEndpoint", common.ConfigurationEndpoint).
		Msg("starting application with provided configuration")

	healthServer := health.PrepareHealthEndpoints(log, common.HealthchecksPort)
	go func() {
		l.Debug().Msg("health endpoints starting")
		if err := healthServer.ListenAndServe(); err != nil {
			l.Fatal().Err(err).Msg("health endpoints have been shut down unexpectedly")
		}
	}()

	if common.DefaultConfigOnStartup {
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
				l.Info().Msg("service has been shut down for configuration change")
			} else {
				l.Fatal().Err(err).Msg("service has been shut down unexpectedly")
			}
		}
	}
}
