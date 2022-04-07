package logs

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

// Initialize prepares logger for first use
func Initialize(pretty bool, level string) {
	desiredLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		desiredLevel = zerolog.InfoLevel
		defer func() {
			logger.Warn().Str("levelName", level).Msg("unknown log level selected, falling back to INFO")
		}()
	}
	zerolog.SetGlobalLevel(desiredLevel)

	var writer io.Writer = os.Stdout
	if pretty {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05",
		}
	}

	logger = zerolog.New(writer).With().Timestamp().Logger()
}

// For returns a logger instance for given component
func For(component string) zerolog.Logger {
	return logger.With().Str("component", component).Logger()
}
