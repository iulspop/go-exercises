package reverse_vowels

import (
	"go-exercises/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	result := ReverseVowels("hello")
	expected := "holle"
	assert.Equals(t, "ReverseVowels()", "Given a string", "Should reverses characters", result, expected)

	result = ReverseVowels("aA")
	expected = "Aa"
	assert.Equals(t, "ReverseVowels()", "Given a string", "Should reverses characters", result, expected)
}
