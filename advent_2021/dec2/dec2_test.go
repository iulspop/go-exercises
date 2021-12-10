package dec2

import (
	"go-exercises/assert"
	"testing"
	"io/ioutil"
)

func TestPosition(t *testing.T) {
	input := `forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2`
	expected := [2]int{15, 10} // Horizontal Position & Depth

	result := Position(input)
	assert.Equals(t, "Position()", "Given", "Should", result, expected)

	input2, _ := ioutil.ReadFile("./puzzle_input.txt")
	expected = [2]int{2199, 786}

	result = Position(string(input2))
	assert.Equals(t, "Position()", "Given", "Should", result, expected)
}