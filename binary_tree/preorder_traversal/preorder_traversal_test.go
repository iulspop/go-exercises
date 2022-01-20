package preorder_traversal

import (
	"go-exercises/assert"
	"testing"
	"go-exercises/binary_tree/utils"
)

func TestPreorderTraversal(t *testing.T) {
	result := PreorderTraversal(utils.CreateTree([]int {4,2,7,1,3,6,9}))
	expected := []int {4,2,1,3,7,6,9}
	assert.Equals(t, "PreorderTraversal()", "", "", result, expected)
}

/*

       4

  2        7

1   3     6   9


4 2 1 3 7 6 9

*/