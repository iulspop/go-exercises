package max_subarray

import (
	"go-exercises/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	result := MaxSubArray([]int {-2, 1, 0})
	expected := 1
	assert.Equals(t, "MaxSubArray()", "", "", result, expected)
}
