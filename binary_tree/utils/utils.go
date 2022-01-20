package utils

import (
	"go-exercises/binary_tree"
	"strconv"
	"strings"
)

func CreateTree(nodes []int) *binary_tree.Node {
	var createNode func(int) *binary_tree.Node
	createNode = func(index int) *binary_tree.Node {
		if index >= len(nodes) { return nil }
		return &binary_tree.Node{Val: nodes[index], Left: createNode(index * 2 + 1), Right: createNode(index * 2 + 2)}
	}
	return createNode(0)
}

func IsSameTree(p *binary_tree.Node, q *binary_tree.Node) bool {
  if p == nil { return q == nil }
  if q == nil { return p == nil }
  if p.Val != q.Val { return false }

  return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func ToString(n *binary_tree.Node) string {
	return "[" + strings.Join(MapToString(LevelOrder(n)), ",") + "]"
}

func LevelOrder(root *binary_tree.Node) []int {
  queue := []*binary_tree.Node {root}
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

func MapToString(nums []int) []string {
	numStrings := []string {}

	for _, num := range nums {
		numStrings = append(numStrings, strconv.Itoa(num))
	}

	return numStrings
}
