package common

import (
	"os"
	"strconv"
	"testing"

	"git.sr.ht/~icikowski/gpts/utils"
	"github.com/stretchr/testify/require"
)

func TestStringPassthrough(t *testing.T) {
	expected := "Lorem ipsum"
	actual, err := stringPassthrough(expected)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestGetFromEnvironment(t *testing.T) {
	testsForStrings := map[string]struct {
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
			valueToSet:                   utils.PointerTo("ABC"),
			expectedVariableDefaultValue: "ABC",
			expectedVariableValue:        "ABC",
		},
		"variable set with value different than default": {
			valueToSet:                   utils.PointerTo("XYZ"),
			expectedVariableDefaultValue: "ABC",
			expectedVariableValue:        "XYZ",
		},
	}

	testsForBooleans := map[string]struct {
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
			valueToSet:                   utils.PointerTo("true"),
			expectedVariableDefaultValue: true,
			expectedVariableValue:        true,
		},
		"variable set with value different than default (case 1)": {
			valueToSet:                   utils.PointerTo("true"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        true,
		},
		"variable set with value different than default (case 2)": {
			valueToSet:                   utils.PointerTo("1"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        true,
		},
		"variable set with malformed value different than default": {
			valueToSet:                   utils.PointerTo("yes"),
			expectedVariableDefaultValue: false,
			expectedVariableValue:        false,
		},
	}

	testsForIntegers := map[string]struct {
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
			valueToSet:                   utils.PointerTo("80"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        80,
		},
		"variable set with value different than default": {
			valueToSet:                   utils.PointerTo("8181"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        8181,
		},
		"variable set with inconvertible value different than default": {
			valueToSet:                   utils.PointerTo("NotANumber"),
			expectedVariableDefaultValue: 80,
			expectedVariableValue:        80,
		},
	}

	t.Run("strings", func(t *testing.T) {
		for name, tc := range testsForStrings {
			t.Run(name, func(t *testing.T) {
				_ = os.Unsetenv("TEST")
				if tc.valueToSet != nil {
					os.Setenv("TEST", string(*tc.valueToSet))
				}

				actualVariable := getFromEnvironment("TEST", tc.expectedVariableDefaultValue, stringPassthrough)
				require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
			})
		}
	})
	t.Run("booleans", func(t *testing.T) {
		for name, tc := range testsForBooleans {
			t.Run(name, func(t *testing.T) {
				_ = os.Unsetenv("TEST")
				if tc.valueToSet != nil {
					os.Setenv("TEST", string(*tc.valueToSet))
				}

				actualVariable := getFromEnvironment("TEST", tc.expectedVariableDefaultValue, strconv.ParseBool)
				require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
			})
		}
	})
	t.Run("integers", func(t *testing.T) {
		for name, tc := range testsForIntegers {
			t.Run(name, func(t *testing.T) {
				_ = os.Unsetenv("TEST")
				if tc.valueToSet != nil {
					os.Setenv("TEST", string(*tc.valueToSet))
				}

				actualVariable := getFromEnvironment("TEST", tc.expectedVariableDefaultValue, strconv.Atoi)

				require.Equal(t, tc.expectedVariableValue, actualVariable, "got value different than expected")
			})
		}
	})
}
