package reverse_string

import (
	"go-exercises/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	result := ReverseString("abcde")
	expected := "edcba"
	assert.Equals(t, "ReverseString()", "Given a string", "Should reverses characters", result, expected)
}
