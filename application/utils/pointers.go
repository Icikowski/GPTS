package utils

// StringToPointer takes string literal and converts it to a pointer
func StringToPointer(in string) *string {
	out := new(string)
	*out = in
	return out
}

// IntToPointer takes int literal and converts it to a pointer
func IntToPointer(in int) *int {
	out := new(int)
	*out = in
	return out
}
