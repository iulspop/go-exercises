package binary_search

import (
	"go-exercises/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	input := []int {1, 2, 3, 8, 9, 30, 40, 100}

	result := BinarySearch(input, 40)
	expected := 6

	assert.Equals(t, "BinarySearch()", "Given slice of integers and search value", "Should return index", result, expected)

	result = BinarySearch(input, 4)
	expected = -1

	assert.Equals(t, "BinarySearch()", "Given slice of integers that do not contain search value", "Should return -1", result, expected)
}
