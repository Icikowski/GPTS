package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/utils"
)

func TestGetFromEnvironment(t *testing.T) {
	tests := map[string]struct {
		expectedVariableName         string
		valueToSet                   *string
		expectedVariableDefaultValue string
		expectedVariableValue        string
	}{
		"variable not set": {
			valueToSet:                   nil,
			expectedVariableDefaultValue: "ABC",
			expectedVariableValue:        "ABC",
		},
		"variable set with value same as default": {
			valueToSet:                   utils.StringToPointer("ABC"),
			expectedVariableDefaultValue: "ABC",
			expectedVariableValue:        "ABC",
		},
		"variable set with value different than default": {
			valueToSet:                   utils.StringToPointer("XYZ"),
			expectedVariableDefaultValue: "ABC",
			expectedVariableValue:        "XYZ",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_ = os.Unsetenv("TEST")
			if tc.valueToSet != nil {
				os.Setenv("TEST", string(*tc.valueToSet))
			}

			actualVariable := getFromEnvironment("TEST", tc.expectedVariableDefaultValue)

			require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
		})
	}
}
