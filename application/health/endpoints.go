package health

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"pkg.icikowski.pl/kubeprobes"
)

// PrepareHealthEndpoints prepares and configures health endpoints
func PrepareHealthEndpoints(log zerolog.Logger, port int) *http.Server {
	log.Debug().
		Int("port", port).
		Msg("preparing readiness & liveness endpoints")

	health, _ := kubeprobes.New(
		kubeprobes.WithLivenessProbes(ApplicationStatus),
		kubeprobes.WithReadinessProbes(ServiceStatus),
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: health,
	}
}
