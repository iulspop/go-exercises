package search_insert_position

import (
	"go-exercises/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	result := searchInsert([]int {1, 3, 5, 6}, 2)
	expected := 1
	assert.Equals(t, "SearchInsert()", "Given a target that cannot be found", "Should return index where insert", result, expected)

	result = searchInsert([]int {1, 3, 5, 6}, 7)
	expected = 4
	assert.Equals(t, "SearchInsert()", "Given a target that cannot be found", "Should return index where insert", result, expected)

	result = searchInsert([]int {1, 3, 5, 6}, 5)
	expected = 2
	assert.Equals(t, "SearchInsert()", "Given a target that is found", "Should return target index", result, expected)
}
