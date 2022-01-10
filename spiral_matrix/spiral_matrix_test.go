package spiral_matrix

import (
	"go-exercises/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	result := SpiralOrder([][]int {{1,2,3},{4,5,6},{7,8,9}})
	expected := []int {1,2,3,6,9,8,7,4,5}
	assert.Equals(t, "SpiralOrder()", "Given a matrix", "Should return elements in clockwise order", result, expected)

	result = SpiralOrder([][]int {{1,2,3,4},{5,6,7,8},{9,10,11,12}})
	expected = []int {1,2,3,4,8,12,11,10,9,5,6,7}
	assert.Equals(t, "SpiralOrder()", "Given a matrix", "Should return elements in clockwise order", result, expected)
}
