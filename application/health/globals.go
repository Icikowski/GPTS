package health

var (
	// ApplicationStatus represents current application liveness
	ApplicationStatus = status{readiness: false}

	// TestServiceStatus represents current test service readiness
	TestServiceStatus = status{readiness: false}
)
