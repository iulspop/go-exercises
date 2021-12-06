package dec2part2

import (
	"go-exercises/assert"
	"io/ioutil"
	"testing"
)

func TestPosition(t *testing.T) {
	input := `forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2`
	expected := [2]int{15, 60} // Horizontal Position & Depth

	result := Position(input)
	assert.Equals(t, "Position()", result, expected)

	input2, _ := ioutil.ReadFile("../dec2/puzzle_input.txt")
	expected = [2]int{2199, 802965}

	result = Position(string(input2))
	assert.Equals(t, "Position()", result, expected)
}