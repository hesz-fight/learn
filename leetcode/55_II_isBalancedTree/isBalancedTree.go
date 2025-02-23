package isbalancedtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans bool

func isBalanceTree(root *TreeNode) bool {
	ans = true
	helper(root)
	return ans
}

// 在计算二叉树深度的时候判断是否是平衡二叉树
func helper(root *TreeNode) int {
	if !ans || root == nil {
		return 0
	}
	l := helper(root.Left)
	r := helper(root.Right)
	if abs(l-r) > 1 {
		ans = false
	}

	return max(l, r)
}

func max(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1
	}
	return x
}
