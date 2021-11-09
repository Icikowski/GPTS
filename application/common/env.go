package common

import (
	"os"
	"strconv"
	"strings"
)

func getStringFromEnvironment(variableName string, fallback string) string {
	value, exists := os.LookupEnv(variableName)
	if !exists {
		return fallback
	}
	return value
}

func getBooleanFromEnvironment(variableName string, fallback bool) bool {
	value, exists := os.LookupEnv(variableName)
	if !exists {
		return fallback
	}
	value = strings.ToLower(value)
	return (value == "true" || value == "1" || value == "yes")
}

func getIntegerFromEnvironment(variableName string, fallback int) int {
	value, exists := os.LookupEnv(variableName)
	if !exists {
		return fallback
	}

	if parsedValue, err := strconv.Atoi(value); err != nil {
		return fallback
	} else {
		return parsedValue
	}
}

var (
	// ServicePort determines the port number on which service will be running (defaults to 80)
	ServicePort = getIntegerFromEnvironment("GPTS_SERVICE_PORT", 80)

	// HealthchecksPort determines the port number on which liveness & readiness endpoints will be running (defaults to 8081)
	HealthchecksPort = getIntegerFromEnvironment("GPTS_HEALTHCHECKS_PORT", 8081)

	// DefaultConfigOnStartup determines if default config should be loaded when application starts
	DefaultConfigOnStartup = getBooleanFromEnvironment("GPTS_DEFAULT_CONFIG_ON_STARTUP", false)

	// PrettyLog determines if pretty logging should be enabled
	PrettyLog = getBooleanFromEnvironment("GPTS_PRETTY_LOG", false)

	// LogLevel determines the level of application log
	LogLevel = getStringFromEnvironment("GPTS_LOG_LEVEL", "info")
)
