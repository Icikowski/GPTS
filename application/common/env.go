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
	// ServicePort determines the port number on which service will be running (defaults to 80)
	ServicePort = getFromEnvironment("GPTS_SERVICE_PORT", "80")

	// HealthchecksPort determines the port number on which liveness & readiness endpoints will be running (defaults to 8081)
	HealthchecksPort = getFromEnvironment("GPTS_HEALTHCHECKS_PORT", "8081")

	// DefaultConfigOnStartup determines if default config should be loaded when application starts
	DefaultConfigOnStartup = getFromEnvironment("GPTS_DEFAULT_CONFIG_ON_STARTUP", "false")

	// LogLevel determines the level of application log
	LogLevel = getFromEnvironment("GPTS_LOG_LEVEL", "info")
)
