package utils

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree"
)

func TestCreateTree(t *testing.T) {
	result := IsSameTree(CreateTree([]int {4,2,7}), &binary_tree.Node{Val: 4, Left: &binary_tree.Node{Val: 2}, Right: &binary_tree.Node{Val: 7}})
	expected := true
	assert.Equals(t, "CreateTree()", "", "", result, expected)

	result2 := ToString(CreateTree([]int {4,2,7,1,3,6,8}))
	expected2 := "[4,2,7,1,3,6,8]"
	assert.Equals(t, "CreateTree()", "", "", result2, expected2)
}

