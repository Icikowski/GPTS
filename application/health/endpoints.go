package health

import (
	"errors"
	"net/http"

	"github.com/heptiolabs/healthcheck"
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
func PrepareHealthEndpoints(port string) *http.Server {
	log.Info().
		Str("port", port).
		Msg("preparing readiness & liveness endpoints")

	health := healthcheck.NewHandler()
	health.AddLivenessCheck("application", applicationLivenessCheck)
	health.AddReadinessCheck("gpts", serviceReadinessCheck)

	return &http.Server{
		Addr:    ":" + port,
		Handler: health,
	}
}
