package roman_numerals

import (
	"go-exercises/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	result := RomanToInt("III")
	expected := 3
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string", "Should add 'I' symbols", result, expected)

	result = RomanToInt("LVIII")
	expected = 58
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string", "Should add 'L', 'V' and 'I' symbols", result, expected)

	result = RomanToInt("MCMXCIV")
	expected = 1994
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string with substractions", "Should calculate substrastions correctly", result, expected)
}
