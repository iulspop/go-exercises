package find_duplicate

import (
	"go-exercises/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	result := FindDuplicate([]string {"a", "b", "c", "d", "c", "e", "f"})
	expected := "c"
	assert.Equals(t, "FindDuplicate()", "Given two slices of integers", "Should return slice of common integers", result, expected)
}
