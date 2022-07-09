package utils

// GetOptionalOrFallback returns pointer to value if value is not empty or fallback pointer otherwise
func GetOptionalOrFallback[T any](value *T, fallback *T) *T {
	if value == nil {
		return fallback
	}
	return value
}
