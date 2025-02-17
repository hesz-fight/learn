package subtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 在roo1中枚举所有的根节点
func isSubStructure(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	return subtree(root1, root2) || subtree(root1.Left, root2) || subtree(root2.Right, root2)
}

func subtree(root1 *TreeNode, root2 *TreeNode) bool {
	// root2为nil或者root1的当前节点值等于root2的当前节点值才能通过当前节点的判断
	if root2 == nil {
		return true
	}
	if root1 == nil || root1.Val != root2.Val {
		return false
	}

	return subtree(root1.Left, root2.Left) && subtree(root1.Right, root2.Right)
}
