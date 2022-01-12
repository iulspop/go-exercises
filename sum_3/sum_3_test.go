package sum_3

import (
	"go-exercises/assert"
	"testing"
)

func TestSum3(t *testing.T) {
	result := Sum3([]int {-1,0,1,2,-1,-4})
	expected := [][]int {{-1,0,1}, {-1,-1,2}}
	assert.Equals(t, "Sum3()", "", "", result, expected)

	result = Sum3([]int {3,0,-2,-1,1,2})
	expected = [][]int {{-2,-1,3},{-2,0,2},{-1,0,1}}
	assert.Equals(t, "Sum3()", "", "", result, expected)
}

func TestIndexNums(t *testing.T) {
	result := indexNums([]int {-1,0,1,2,-1,-4})
	expected := numsIndexes{
		map[int]int {
			1: 1,
			2: 1,
		},
		map[int]int {
			-1: 2,
			-4: 1,
		},
		map[int]int {
			0: 1,
		},
	}

	assert.Equals(t, "IndexNums()", "", "", result, expected)
}

