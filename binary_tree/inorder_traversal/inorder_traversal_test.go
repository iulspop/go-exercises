package inorder_traversal

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree/utils"
)

func TestInorderTraversal(t *testing.T) {
	result := InorderTraversal(utils.CreateTree([]int {4,2,7,1,3,6,9}))
	expected := []int {1,2,3,4,6,7,9}
	assert.Equals(t, "InorderTraversal()", "", "", result, expected)
}

func TestInorderTraversalWithStack(t *testing.T) {
	result := InorderTraversalWithStack(utils.CreateTree([]int {4,2,7,1,3,6,9}))
	expected := []int {1,2,3,4,6,7,9}
	assert.Equals(t, "InorderTraversalWithStack()", "", "", result, expected)
}

/*

       4

  2        7

1   3     6   9


1,2,3,4,6,7,9

*/