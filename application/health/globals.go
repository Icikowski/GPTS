package health

var (
	// ApplicationStatus represents current application liveness
	ApplicationStatus = status{readiness: false}

	// ServiceStatus represents current test service readiness
	ServiceStatus = status{readiness: false}
)
