package invert_binary_tree

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree/utils"
)

func TestInvertTree(t *testing.T) {
	result := utils.ToString(InvertTree(utils.CreateTree([]int {4,2,7,1,3,6,9})))
	expected := "[4,7,2,9,6,3,1]"
	assert.Equals(t, "InvertTree()", "", "", result, expected)
}
