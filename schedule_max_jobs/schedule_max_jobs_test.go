package schedule_max_jobs

import (
	"go-exercises/assert"
	"testing"
)

func TestScheduleMaxJobs(t *testing.T) {
	input := [][2]int{{4, 12}, {2, 9}, {10, 15}, {6, 15}, {14, 34}, {16, 20}, {21, 30}, {22, 30}, {28, 46}, {32, 48}}

	result := ScheduleMaxJobs(input)
	expected := [][2]int{{2, 9}, {10, 15}, {16, 20}, {21, 30}, {32, 48}}

	assert.Equals(t, "ScheduleMaxJobs()", "Given set of intervals", "Should return largest set of non-overlapping intervals", result, expected)
}

func TestSortIntervalsByEnd(t *testing.T) {
	input := [][2]int{{32, 48}, {4, 12}, {2, 9}, {14, 34}, {10, 15}, {6, 15}}

	result := SortIntervalsByEnd(input)
	expected := [][2]int{{2, 9}, {4, 12}, {10, 15}, {6, 15}, {14, 34}, {32, 48}}

	assert.Equals(t, "SortIntervalsByEnd()", "Given set of intervals", "Should sort them by the end time", result, expected)
}

