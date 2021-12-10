package bubble_sort

import (
	"go-exercises/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	result := BubbleSort([]int {4, 2, 7, 1, 3})
	expected := []int {1, 2, 3, 4, 7}
	assert.Equals(t, "BubbleSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)

	result = BubbleSort([]int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	expected = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equals(t, "BubbleSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)
}
