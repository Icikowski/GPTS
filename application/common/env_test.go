package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/utils"
)

func TestGetStringFromEnvironment(t *testing.T) {
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

			actualVariable := getStringFromEnvironment("TEST", tc.expectedVariableDefaultValue)

			require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
		})
	}
}

func TestGetBooleanFromEnvironment(t *testing.T) {
	tests := map[string]struct {
		expectedVariableName         string
		valueToSet                   *string
		expectedVariableDefaultValue bool
		expectedVariableValue        bool
	}{
		"variable not set": {
			valueToSet:                   nil,
			expectedVariableDefaultValue: true,
			expectedVariableValue:        true,
		},
		"variable set with value same as default": {
			valueToSet:                   utils.StringToPointer("true"),
			expectedVariableDefaultValue: true,
			expectedVariableValue:        true,
		},
		"variable set with value different than default (case 1)": {
			valueToSet:                   utils.StringToPointer("true"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        true,
		},
		"variable set with value different than default (case 2)": {
			valueToSet:                   utils.StringToPointer("yes"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        true,
		},
		"variable set with value different than default (case 3)": {
			valueToSet:                   utils.StringToPointer("1"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_ = os.Unsetenv("TEST")
			if tc.valueToSet != nil {
				os.Setenv("TEST", string(*tc.valueToSet))
			}

			actualVariable := getBooleanFromEnvironment("TEST", tc.expectedVariableDefaultValue)

			require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
		})
	}
}

func TestGetIntegerFromEnvironment(t *testing.T) {
	tests := map[string]struct {
		expectedVariableName         string
		valueToSet                   *string
		expectedVariableDefaultValue int
		expectedVariableValue        int
	}{
		"variable not set": {
			valueToSet:                   nil,
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        80,
		},
		"variable set with value same as default": {
			valueToSet:                   utils.StringToPointer("80"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        80,
		},
		"variable set with value different than default": {
			valueToSet:                   utils.StringToPointer("8181"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        8181,
		},
		"variable set with inconvertible value different than default": {
			valueToSet:                   utils.StringToPointer("NotANumber"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        80,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_ = os.Unsetenv("TEST")
			if tc.valueToSet != nil {
				os.Setenv("TEST", string(*tc.valueToSet))
			}

			actualVariable := getIntegerFromEnvironment("TEST", tc.expectedVariableDefaultValue)

			require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
		})
	}
}
