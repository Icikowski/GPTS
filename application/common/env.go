package common

import "strconv"

var (
	// ServicePort determines the port number on which service will be running (defaults to 80)
	ServicePort int = getFromEnvironment("GPTS_SERVICE_PORT", 80, strconv.Atoi)

	// HealthchecksPort determines the port number on which liveness & readiness endpoints will be running (defaults to 8081)
	HealthchecksPort int = getFromEnvironment("GPTS_HEALTHCHECKS_PORT", 8081, strconv.Atoi)

	// ConfigurationEndpoint determines the path of the configuration endpoint (defaults to /config)
	ConfigurationEndpoint string = getFromEnvironment("GPTS_CONFIG_ENDPOINT", "/config", stringPassthrough)

	// DefaultConfigOnStartup determines if default config should be loaded when application starts (defaults to false)
	DefaultConfigOnStartup bool = getFromEnvironment("GPTS_DEFAULT_CONFIG_ON_STARTUP", false, strconv.ParseBool)

	// PrettyLog determines if pretty logging should be enabled (defaults to false)
	PrettyLog bool = getFromEnvironment("GPTS_PRETTY_LOG", false, strconv.ParseBool)

	// LogLevel determines the level of application log (defaults to "info")
	LogLevel string = getFromEnvironment("GPTS_LOG_LEVEL", "info", stringPassthrough)
)
