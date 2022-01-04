package non_duplicated_letter

import (
	"go-exercises/assert"
	"testing"
)

func TestNonDuplicatedLetter(t *testing.T) {
	result := NonDuplicatedLetter("minimum" )
	expected := "n"
	assert.Equals(t, "NonDuplicatedLetter()", "Given a string", "Should first non-duplicated char", result, expected)
}
