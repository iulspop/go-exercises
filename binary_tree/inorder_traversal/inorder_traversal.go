package inorder_traversal

import "go-exercises/binary_tree"

func InorderTraversal(root *binary_tree.Node) (slice []int) {
	if root == nil { return }

	slice = append(slice, InorderTraversal(root.Left)...)
	slice = append(slice, root.Val)
	slice = append(slice, InorderTraversal(root.Right)...)

	return
}

func InorderTraversalWithStack(root *binary_tree.Node) []int {
	stack := []*binary_tree.Node {root}
	slice := []int {}

	left := true
	for len(stack) != 0 {
		node := stack[len(stack) - 1]

	  if left && node.Left != nil { stack = append(stack, node.Left); continue }

		node = stack[len(stack) - 1]
		slice = append(slice, node.Val)
		stack = stack[:len(stack) - 1]
		left = false

		if node.Right != nil { stack = append(stack, node.Right); left = true }
	}

	return slice
}
