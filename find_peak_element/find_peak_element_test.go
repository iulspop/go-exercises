package find_peak_element

import (
	"go-exercises/assert"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	result := FindPeakElement([]int {0})
	expected := 0
	assert.Equals(t, "FindPeakElement()", "", "", result, expected)

	result = FindPeakElement([]int {1, 2})
	expected = 1
	assert.Equals(t, "FindPeakElement()", "", "", result, expected)

	result = FindPeakElement([]int {2, 1})
	expected = 0
	assert.Equals(t, "FindPeakElement()", "", "", result, expected)

	result = FindPeakElement([]int {1,2,1,3,5,6,4})
	expected = 5
	assert.Equals(t, "FindPeakElement()", "", "", result, expected)
}
