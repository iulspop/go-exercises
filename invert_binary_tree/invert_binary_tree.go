package invert_binary_tree

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
  if (root == nil) { return root }
  if (root.Left == nil && root.Right == nil) { return root }

  root.Left, root.Right = root.Right, root.Left
  InvertTree(root.Left); InvertTree(root.Right);

  return root
}

/*

implicit requirements:
- no guaranteed to be balanced (one side can be nil)

base case:
  both left & right are nil
  
definition:
  right & left are switched && right subtree inverted && left subtree inverted

  return root

*/