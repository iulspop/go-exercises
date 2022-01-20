package preorder_traversal

import "go-exercises/binary_tree"

func PreorderTraversal(root *binary_tree.Node) (slice []int) {
	if root == nil { return }

	slice = append(slice, root.Val)
	slice = append(slice, PreorderTraversal(root.Left)...)
	slice = append(slice, PreorderTraversal(root.Right)...)

	return
}

func PreorderTraversalWithStack(root *binary_tree.Node) []int {
	stack := []*binary_tree.Node {root}
	slice := []int {}

	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		slice = append(slice, node.Val)

		stack = stack[:len(stack) - 1]

		if node.Right != nil { stack = append(stack, node.Right) }
		if node.Left  != nil { stack = append(stack, node.Left) }
	}

	return slice
}

/*




*/