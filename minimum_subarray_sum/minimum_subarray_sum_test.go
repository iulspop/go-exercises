package minimum_subarray_sum

import (
	"go-exercises/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	result := minSubArrayLen(7, []int {2,3,1,2,4,3})
	expected := 2
	assert.Equals(t, "minSubArrayLen()", "", "", result, expected)

	result = minSubArrayLen(4, []int {1, 4, 4})
	expected = 1
	assert.Equals(t, "minSubArrayLen()", "", "", result, expected)

	result = minSubArrayLen(11, []int {1, 1, 1, 1, 1, 1, 1})
	expected = 0
	assert.Equals(t, "minSubArrayLen()", "", "", result, expected)
}