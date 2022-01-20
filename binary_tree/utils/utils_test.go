package utils

import (
	"go-exercises/assert"
	"testing"
)

func TestCreateTree(t *testing.T) {
	result := IsSameTree(CreateTree([]int {4,2,7}), &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}})
	expected := true
	assert.Equals(t, "CreateTree()", "", "", result, expected)

	result2 := ToString(CreateTree([]int {4,2,7,1,3,6,8}))
	expected2 := "[4,2,7,1,3,6,8]"
	assert.Equals(t, "CreateTree()", "", "", result2, expected2)
}

