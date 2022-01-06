package combos

import (
	"go-exercises/assert"
	"testing"
)

func TestCombos(t *testing.T) {
	result := Combos([]int {5}, map[string]bool {})
	expected := [][]int {{5}, {1, 4}, {1, 1, 3}, {1, 1, 1, 2}, {1, 1, 1, 1, 1},  {1, 2, 2}, {2, 3}}
	assert.Equals(t, "Combos()", "Given a number", "Should return all possible sets that sum to that number", result, expected)

	// result = Combos([]int {1,1,1,1,1})
	// expected = [][]int {{1,1,1,1,1}}
	// assert.Equals(t, "Combos()", "Given a number", "Should return all possible sets that sum to that number", result, expected)

	// result = Combos([]int {1,1,1,2})
	// expected = [][]int {{1,1,1,1,1}}
	// assert.Equals(t, "Combos()", "Given a number", "Should return all possible sets that sum to that number", result, expected)
}

func TestRemove(t *testing.T)  {
	result := remove([]int {1, 1, 3}, 2)
	expected := []int {1, 1}
	assert.Equals(t, "revome()", "Given a number", "Should return all possible sets that sum to that number", result, expected)
}