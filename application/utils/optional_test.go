package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOptionalOrFallback(t *testing.T) {
	someValue, someFallback := PointerTo("someValue"), PointerTo("someFallback")
	tests := map[string]struct {
		value         *string
		fallback      *string
		expectedValue *string
	}{
		"value is not nil": {
			value:         someValue,
			fallback:      someFallback,
			expectedValue: someValue,
		},
		"value is nil": {
			value:         nil,
			fallback:      someFallback,
			expectedValue: someFallback,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectedValue, GetOptionalOrFallback(tc.value, tc.fallback))
		})
	}
}
