package postorder_traversal

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree/utils"
)

func TestPostOrderTraversal(t *testing.T) {
	result := PostOrderTraversal(utils.CreateTree([]int {4,2,7,1,3,6,9}))
	expected := []int {1,3,2,6,9,7,4}
	assert.Equals(t, "PostOrderTraversal()", "", "", result, expected)
}

func TestPostOrderTraversalWithStack(t *testing.T) {
	result := PostOrderTraversalWithStack(utils.CreateTree([]int {4,2,7,1,3,6,9}))
	expected := []int {1,3,2,6,9,7,4}
	assert.Equals(t, "PostOrderTraversalWithStack()", "", "", result, expected)
}

/*

       4

  2        7

1   3     6   9


1,3,2,6,9,7,4




4 2 3

*/