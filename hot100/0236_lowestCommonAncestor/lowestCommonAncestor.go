package lowestcommonancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 返回p、q节点的最近公共祖先
// 当前节点的返回值根据左子树和右子树的情况不同来判断
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	lr := lowestCommonAncestor(root.Left, p, q)
	rr := lowestCommonAncestor(root.Right, p, q)
	if lr == nil {
		return rr
	}
	if rr == nil {
		return lr
	}
	return root
}
