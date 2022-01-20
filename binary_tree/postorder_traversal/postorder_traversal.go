package postorder_traversal

import "go-exercises/binary_tree"

func PostOrderTraversal(root *binary_tree.Node) (slice []int) {
	if root == nil { return }

	slice = append(slice, PostOrderTraversal(root.Left)...)
	slice = append(slice, PostOrderTraversal(root.Right)...)
	slice = append(slice, root.Val)

	return
}

func PostOrderTraversalWithStack(root *binary_tree.Node) []int {
	stack := []*binary_tree.Node {root}
	slice := []int {}

	left := true
	for len(stack) != 0 {
		node := stack[len(stack) - 1]

	  if left && node.Left != nil { stack = append(stack, node.Left); continue }
		if !left && node.Right != nil { stack = append(stack, node.Right); left = true }

		node = stack[len(stack) - 1]
		slice = append(slice, node.Val)
		stack = stack[:len(stack) - 1]
		left = false
	}

	return slice
}