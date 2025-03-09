package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 先序遍历
	if pre != nil {
		pre.Left = nil
		pre.Right = root
	}
	pre = root
	right := pre.Right
	flatten(root.Left)
	flatten(right)
}
