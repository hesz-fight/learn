package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return recur(root.Left, root.Right)
}

func recur(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	return recur(r1.Left, r2.Right) && recur(r1.Right, r2.Left)
}
