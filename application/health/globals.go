package health

import (
	"pkg.icikowski.pl/kubeprobes"
)

var (
	// ApplicationStatus represents current application liveness
	ApplicationStatus, _ = kubeprobes.NewManualProbe("app")

	// ServiceStatus represents current test service readiness
	ServiceStatus, _ = kubeprobes.NewManualProbe("svc")
)
