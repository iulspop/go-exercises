package is_symmetric

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
  return isMirror(root.Left, root.Right)
}

func isMirror(leftRoot *TreeNode, rightRoot *TreeNode) bool {
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
