package selection_sort

import (
	"go-exercises/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	result := SelectionSort([]int {4, 2, 7, 1, 3})
	expected := []int {1, 2, 3, 4, 7}
	assert.Equals(t, "SelectionSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)

	result = SelectionSort([]int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	expected = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equals(t, "SelectionSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)
}
