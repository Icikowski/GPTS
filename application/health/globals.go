package health

import (
	"github.com/Icikowski/kubeprobes"
)

var (
	// ApplicationStatus represents current application liveness
	ApplicationStatus = kubeprobes.NewStatefulProbe()

	// ServiceStatus represents current test service readiness
	ServiceStatus = kubeprobes.NewStatefulProbe()
)
