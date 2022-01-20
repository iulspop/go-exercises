package is_symmetric

import "go-exercises/binary_tree"

func IsSymmetric(root *binary_tree.Node) bool {
  return isMirror(root.Left, root.Right)
}

func isMirror(leftRoot *binary_tree.Node, rightRoot *binary_tree.Node) bool {
  if leftRoot == nil && rightRoot == nil { return true }
  if leftRoot == nil || rightRoot == nil { return false }

  return leftRoot.Val == rightRoot.Val && isMirror(leftRoot.Left, rightRoot.Right) && isMirror(leftRoot.Right, rightRoot.Left)
}

/*

base case:

definition:
  if both roots nil { return true }
  if right || left nil { return false }

  leftRoot and rightRoot are mirror if
    they have same value && mirror(leftRoot.left, rightRoot.right) && mirror(leftRoot.right, rightRoot.left)

*/
