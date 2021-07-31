package common

import "os"

func getFromEnvironment(variableName string, fallback string) string {
	value, exists := os.LookupEnv(variableName)
	if !exists {
		value = fallback
	}
	return value
}

var (
	// ServerPort determines the port number on which service will be running (defaults to 80)
	ServerPort = getFromEnvironment("GPTS_SERVER_PORT", "80")

	// HealthcheckPort determines the port number on which liveness & readiness endpoints will be running (defaults to 8000)
	HealthcheckPort = getFromEnvironment("GPTS_HEALTHCHECK_PORT", "8000")

	// DefaultConfigOnStartup determines if default config should be loaded when application starts
	DefaultConfigOnStartup = getFromEnvironment("GPTS_DEFAULT_CONFIG_ON_STARTUP", "false")
)
