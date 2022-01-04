package missing_letter

import (
	"go-exercises/assert"
	"testing"
)

func TestMissingLetter(t *testing.T) {
	result := MissingLetter("the quick brown box jumps over a lazy dog" )
	expected := "f"
	assert.Equals(t, "MissingLetter()", "Given a phrase containing all letters of alphabet except one", "Should return missing letter", result, expected)
}
