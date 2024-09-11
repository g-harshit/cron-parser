package lib

import (
	"strconv"
	"strings"
)

// ConvertIntsToStrings converts a slice of integers to a slice of strings.
func ConvertIntsToStrings(ints []int) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strs
}

func ConvertIntsToSingleString(ints []int) string {
	return strings.Join(ConvertIntsToStrings(ints), " ")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
