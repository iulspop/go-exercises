package is_subsequence

import (
	"go-exercises/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	result := IsSubsequence("abc", "ahbgdc")
	expected := true
	assert.Equals(t, "IsSubsequence()", "", "", result, expected)

	result = IsSubsequence("axc", "ahbgdc")
	expected = false
	assert.Equals(t, "IsSubsequence()", "", "", result, expected)
}
