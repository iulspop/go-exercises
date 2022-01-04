package intersect

import (
	"go-exercises/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	result := Intersect([]int {1, 2, 3, 4, 5}, []int {0, 2, 4, 6, 9})
	expected := []int {2, 4}
	assert.Equals(t, "Intersect()", "Given two slices of integers", "Should return slice of common integers", result, expected)
}
