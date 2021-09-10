package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var mainLog zerolog.Logger = zerolog.New(zerolog.ConsoleWriter{
	NoColor:    false,
	Out:        os.Stdout,
	TimeFormat: zerolog.TimeFormatUnix,
})

// ForComponent creates sublogger instance for given component
func ForComponent(component string) zerolog.Logger {
	return mainLog.With().
		Str("component", component).
		Logger()
}
