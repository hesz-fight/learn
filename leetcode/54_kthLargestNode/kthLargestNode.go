package kthlargestnode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ck int
var ans *int

func kthlargestnode(root *TreeNode, k int) int {
	ck = k
	cnt := 0
	helper(root, &cnt)
	return *ans
}

func helper(root *TreeNode, cnt *int) {
	if root == nil {
		return
	}

	helper(root.Right, cnt)
	if ans != nil {
		return
	}
	*cnt++
	if *cnt == ck {
		ans = &root.Val
		return
	}
	helper(root.Left, cnt)
}
