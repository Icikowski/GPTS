package health

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const testPort = "30100"

func TestPrepareHealthEndpoints(t *testing.T) {
	tests := map[string]struct {
		applicationStatus       bool
		serviceStatus           bool
		expectedLivenessStatus  int
		expectedReadinessStatus int
	}{
		"everything down": {
			applicationStatus:       false,
			serviceStatus:           false,
			expectedLivenessStatus:  http.StatusServiceUnavailable,
			expectedReadinessStatus: http.StatusServiceUnavailable,
		},
		"application down and service up": {
			applicationStatus:       false,
			serviceStatus:           true,
			expectedLivenessStatus:  http.StatusServiceUnavailable,
			expectedReadinessStatus: http.StatusServiceUnavailable,
		},
		"application up and service down": {
			applicationStatus:       true,
			serviceStatus:           false,
			expectedLivenessStatus:  http.StatusOK,
			expectedReadinessStatus: http.StatusServiceUnavailable,
		},
		"everything up": {
			applicationStatus:       true,
			serviceStatus:           true,
			expectedLivenessStatus:  http.StatusOK,
			expectedReadinessStatus: http.StatusOK,
		},
	}

	server := PrepareHealthEndpoints(testPort)
	go func() {
		server.ListenAndServe()
	}()
	defer server.Shutdown(context.Background())

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ApplicationStatus.SetStatus(tc.applicationStatus)
			ServiceStatus.SetStatus(tc.serviceStatus)

			liveness, _ := http.Get("http://localhost:" + testPort + "/live")
			readiness, _ := http.Get("http://localhost:" + testPort + "/ready")

			require.Equal(t, tc.expectedLivenessStatus, liveness.StatusCode, "unexpected liveness status")
			require.Equal(t, tc.expectedReadinessStatus, readiness.StatusCode, "unexpected readiness status")
		})
	}
}
