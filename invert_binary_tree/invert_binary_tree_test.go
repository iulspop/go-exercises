package invert_binary_tree

import (
	"go-exercises/assert"
	"testing"
	"strings"
	"strconv"
)

func TestInvertTree(t *testing.T) {
	// result := InvertTree(createTree)
	// expected :=
	// assert.Equals(t, "InvertTree()", "", "", result, expected)
}

func TestCreateTree(t *testing.T) {
	result := isSameTree(createTree([]int {4,2,7}), &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}})
	expected := true
	assert.Equals(t, "createTree()", "", "", result, expected)

	result2 := toString(createTree([]int {4,2,7,1,3,6,8}))
	expected2 := "[4,2,7,1,3,6,8]"
	assert.Equals(t, "createTree()", "", "", result2, expected2)
}

func createTree(nodes []int) *TreeNode {
	var createNode func(int) *TreeNode
	createNode = func(index int) *TreeNode {
		if index >= len(nodes) { return nil }
		return &TreeNode{Val: nodes[index], Left: createNode(index * 2 + 1), Right: createNode(index * 2 + 2)}
	}
	return createNode(0)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
  if p == nil { return q == nil }
  if q == nil { return p == nil }
  if p.Val != q.Val { return false }

  return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func toString(n *TreeNode) string {
	return "[" + strings.Join(mapToString(levelOrder(n)), ",") + "]"
}

func levelOrder(root *TreeNode) []int {
  queue := []*TreeNode {root}
  order := []int {}

  for len(queue) != 0 {
    node := queue[0]
    order = append(order, node.Val)

    queue = queue[1:]

		if node.Left != nil { queue = append(queue, node.Left) }
		if node.Right != nil { queue = append(queue, node.Right) }
  }

  return order
}

func mapToString(nums []int) []string {
	numStrings := []string {}

	for _, num := range nums {
		numStrings = append(numStrings, strconv.Itoa(num))
	}

	return numStrings
}
