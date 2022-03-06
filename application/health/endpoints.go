package health

import (
	"fmt"
	"net/http"

	"github.com/Icikowski/kubeprobes"
	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
)

// PrepareHealthEndpoints prepares and configures health endpoints
func PrepareHealthEndpoints(log zerolog.Logger, port int) *http.Server {
	l := log.With().Str(common.ComponentField, common.ComponentHealth).Logger()
	l.Debug().
		Int("port", port).
		Msg("preparing readiness & liveness endpoints")

	health := kubeprobes.NewKubeprobes(
		kubeprobes.WithLivenessProbes(ApplicationStatus.GetProbeFunction()),
		kubeprobes.WithReadinessProbes(ServiceStatus.GetProbeFunction()),
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: health,
	}
}
