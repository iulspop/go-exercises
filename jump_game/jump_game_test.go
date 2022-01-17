package jump_game

import (
	"go-exercises/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	result := CanJump([]int {3,2,1,0,4})
	expected := false
	assert.Equals(t, "CanJump()", "", "", result, expected)

	result = CanJump([]int {2,3,1,1,4})
	expected = true
	assert.Equals(t, "CanJump()", "", "", result, expected)

	result = CanJump([]int {0})
	expected = true
	assert.Equals(t, "CanJump()", "", "", result, expected)
}
