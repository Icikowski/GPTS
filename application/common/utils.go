package common

import (
	"os"
)

func getFromEnvironment[T any](key string, fallback T, parser func(string) (T, error)) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	if parsedValue, err := parser(value); err == nil {
		return parsedValue
	}
	return fallback
}

func stringPassthrough(value string) (string, error) {
	return value, nil
}
