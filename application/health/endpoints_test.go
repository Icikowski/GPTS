package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

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

	handler := PrepareHealthEndpoints(zerolog.Nop(), 80).Handler
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	client := testServer.Client()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.applicationStatus {
				ApplicationStatus.Pass()
			} else {
				ApplicationStatus.Fail()
			}

			if tc.serviceStatus {
				ServiceStatus.Pass()
			} else {
				ServiceStatus.Fail()
			}

			liveness, _ := client.Get(testServer.URL + "/live")
			readiness, _ := client.Get(testServer.URL + "/ready")

			require.Equal(t, tc.expectedLivenessStatus, liveness.StatusCode, "unexpected liveness status")
			require.Equal(t, tc.expectedReadinessStatus, readiness.StatusCode, "unexpected readiness status")
		})
	}
}
