package roman_numerals

import (
	"go-exercises/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	result := RomanToInt("III")
	expected := 3
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string", "Should reverses characters", result, expected)

	result = RomanToInt("LVIII")
	expected = 58
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string", "Should reverses characters", result, expected)

	result = RomanToInt("MCMXCIV")
	expected = 1994
	assert.Equals(t, "RomanNumerals()", "Given a roman numeral string where ", "Should reverses characters", result, expected)
}
