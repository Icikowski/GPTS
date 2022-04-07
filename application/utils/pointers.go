package utils

// PointerTo takes some literal and returns a pointer
func PointerTo[T any](in T) *T {
	return &in
}
