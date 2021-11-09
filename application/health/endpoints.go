package health

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
)

func applicationLivenessCheck() error {
	if !ApplicationStatus.GetStatus() {
		return errors.New("DOWN")
	}
	return nil
}

func serviceReadinessCheck() error {
	if !ServiceStatus.GetStatus() {
		return errors.New("DOWN")
	}
	return nil
}

// PrepareHealthEndpoints prepares and configures health endpoints
func PrepareHealthEndpoints(log zerolog.Logger, port int) *http.Server {
	l := log.With().Str(common.ComponentField, common.ComponentHealth).Logger()
	l.Debug().
		Int("port", port).
		Msg("preparing readiness & liveness endpoints")

	health := healthcheck.NewHandler()
	health.AddLivenessCheck("application", applicationLivenessCheck)
	health.AddReadinessCheck("gpts", serviceReadinessCheck)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: health,
	}
}
