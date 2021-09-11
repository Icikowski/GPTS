package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToPointer(t *testing.T) {
	original := "The quick brown fox jumped over the lazy dog"
	pointer := StringToPointer(original)

	require.EqualValues(t, original, *pointer, "values of input string and produced pointer should be equal")
}

func TestIntToPointer(t *testing.T) {
	original := int(12345)
	pointer := IntToPointer(original)

	require.EqualValues(t, original, *pointer, "values of input integer and produced pointer should be equal")
}
