package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
)

var mainLog zerolog.Logger = zerolog.New(os.Stdout).
	With().
	Timestamp().
	Logger()

var logLevels = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
	"trace": zerolog.TraceLevel,
}

// InitializeLog prepares log component for first use
func InitializeLog() {
	zerolog.SetGlobalLevel(logLevels[strings.ToLower(common.LogLevel)])
}

// ForComponent creates sublogger instance for given component
func ForComponent(component string) zerolog.Logger {
	return mainLog.With().
		Str("component", component).
		Logger()
}
