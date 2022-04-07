package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointerTo(t *testing.T) {
	tests := map[string]any{
		"pointer to integer": 50,
		"pointer to string":  "Lorem ipsum",
		"pointer to bool":    true,
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			actual := PointerTo(tc)
			require.Equal(t, tc, *actual, "got value different than expected")
		})
	}
}
