package kthlargestnode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var k0 int
var ans *int

func kthlargestnode(root *TreeNode, k int) int {
	k0 = k
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
	if *cnt == k0 {
		ans = &root.Val
		return
	}
	helper(root.Left, cnt)
}
