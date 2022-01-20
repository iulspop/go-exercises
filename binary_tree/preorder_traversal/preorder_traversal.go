package preorder_traversal

import "go-exercises/binary_tree"

func PreorderTraversal(root *binary_tree.Node) (slice []int) {
	if root == nil { return }

	slice = append(slice, root.Val)
	slice = append(slice, PreorderTraversal(root.Left)...)
	slice = append(slice, PreorderTraversal(root.Right)...)

	return
}
