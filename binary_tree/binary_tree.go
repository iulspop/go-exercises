package binary_tree

import (
	"strings"
	"strconv"
)
type Node struct {
	Val int
	Left *Node
	Right *Node
}

func (n *Node) String() string {
	return toString(n)
}

func toString(n *Node) string {
	return "[" + strings.Join(mapToString(levelOrder(n)), ",") + "]"
}

func levelOrder(root *Node) []int {
  queue := []*Node {root}
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
