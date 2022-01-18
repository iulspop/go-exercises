package minimum_path_sum

import (
	"go-exercises/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	result := MinPathSum([][]int {{1,3,1},{1,5,1},{4,2,1}})
	expected := 7
	assert.Equals(t, "MinPathSum()", "", "", result, expected)

}
