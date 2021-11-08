package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var mainLog zerolog.Logger

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
func InitializeLog(pretty, level string) {
	zerolog.SetGlobalLevel(logLevels[strings.ToLower(level)])
	if pretty == "true" {
		mainLog = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05",
		}).With().Timestamp().Logger()
	} else {
		mainLog = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}

// GetLogger returns a logger instance
func GetLogger() zerolog.Logger {
	return mainLog
}
