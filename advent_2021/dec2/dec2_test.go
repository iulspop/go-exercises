package dec2

import (
	"go-exercises/assert"
	"testing"
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
	assert.Equals(t, "ScheduleMaxJobs()", result, expected)
}