package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(recur(root))
}

// 函数有两个返回值
// 第一个是取根节点的最大返回值，第二个是不取根节点的最大返回值

func recur(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	y1, n1 := recur(root.Left)
	y2, n2 := recur(root.Right)
	// 取根节点时，不能选两个子节点，因此是n1 + n2 + root.Val
	// 不取根节点时，左右子节点可选也可以不选
	return n1 + n2 + root.Val, max(y1, n1) + max(y2, n2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
