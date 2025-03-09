package isvalidbst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre = math.MinInt
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	// 更新上一个中序遍历的值
	pre = root.Val
	return isValidBST(root.Right)
}
