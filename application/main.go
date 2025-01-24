package main

import (
	"flag"
	"net/http"
	"runtime"

	"git.sr.ht/~icikowski/gpts/common"
	"git.sr.ht/~icikowski/gpts/config"
	"git.sr.ht/~icikowski/gpts/health"
	"git.sr.ht/~icikowski/gpts/logs"
	"git.sr.ht/~icikowski/gpts/service"
)

var version = common.BuildValueUnknown
var gitCommit = common.BuildValueUnknown
var binaryType = common.BuildValueUnknown

var logFactory *logs.LoggerFactory

func init() {
	flag.IntVar(&common.ServicePort, "service-port", common.ServicePort, "Port on which the service will be running")
	flag.IntVar(&common.HealthchecksPort, "health-port", common.HealthchecksPort, "Port on which the healthchecks will be running")
	flag.StringVar(&common.ConfigurationEndpoint, "config-endpoint", common.ConfigurationEndpoint, "Path of the configuration endpoint")
	flag.BoolVar(&common.DefaultConfigOnStartup, "default-config", common.DefaultConfigOnStartup, "Enables loading the default configuration on startup")
	flag.BoolVar(&common.PrettyLog, "pretty-log", common.PrettyLog, "Enables the pretty logger")
	flag.StringVar(&common.LogLevel, "log-level", common.LogLevel, "Global log level; one of [debug, info, warn, error, fatal, panic, trace]")
	flag.Parse()

	logFactory = logs.NewFactory(common.PrettyLog, common.LogLevel)
}

func main() {
	log := logFactory.For(common.ComponentCLI)

	log.Info().
		Str("version", version).
		Str("gitCommit", gitCommit).
		Str("binaryType", binaryType).
		Str("goVersion", runtime.Version()).
		Msg("version information")

	log.Info().
		Int("servicePort", common.ServicePort).
		Int("healthchecksPort", common.HealthchecksPort).
		Str("configurationEndpoint", common.ConfigurationEndpoint).
		Msg("configuration applied")

	healthServer := health.PrepareHealthEndpoints(
		logFactory.For(common.ComponentHealth),
		common.HealthchecksPort,
	)
	go func() {
		log.Debug().Msg("health endpoints starting")
		if err := healthServer.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("health endpoints have been shut down unexpectedly")
		}
	}()

	if common.DefaultConfigOnStartup {
		log.Info().Msg("loading default configuration")
		config.CurrentConfiguration.SetDefaultConfiguration(logFactory.For(common.ComponentConfig))
	}

	log.Debug().Msg("marking application liveness as UP")
	health.ApplicationStatus.Pass()

	for {
		service.ExpectingShutdown = false
		server := service.PrepareServer(logFactory.For(common.ComponentService), common.ServicePort)
		health.ServiceStatus.Pass()
		if err := server.ListenAndServe(); err != nil {
			if service.ExpectingShutdown && err == http.ErrServerClosed {
				log.Info().Msg("service has been shut down for configuration change")
			} else {
				log.Fatal().Err(err).Msg("service has been shut down unexpectedly")
			}
		}
	}
}
