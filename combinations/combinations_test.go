package combinations

import (
	"go-exercises/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	result := Combine(4, 2)
	expected := [][]int {{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}}
	assert.Equals(t, "Combine()", "", "", result, expected)
}
