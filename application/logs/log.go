package logs

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// LoggerFactory is a factory capable of preparing Logger objects
type LoggerFactory struct {
	log zerolog.Logger
}

// NewFactory prepares the logger factory for first use
func NewFactory(pretty bool, level string) *LoggerFactory {
	var log zerolog.Logger
	desiredLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		desiredLevel = zerolog.InfoLevel
		defer func() {
			log.Warn().Str("levelName", level).Msg("unknown log level selected, falling back to INFO")
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

	log = zerolog.New(writer).Level(desiredLevel).With().Timestamp().Logger()

	return &LoggerFactory{
		log: log,
	}
}

// For returns a logger instance for given component
func (lf *LoggerFactory) For(component string) zerolog.Logger {
	return lf.log.With().Str("component", component).Logger()
}
