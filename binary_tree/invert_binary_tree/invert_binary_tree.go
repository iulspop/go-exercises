package invert_binary_tree

import "go-exercises/binary_tree"

func InvertTree(root *binary_tree.Node) *binary_tree.Node {
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