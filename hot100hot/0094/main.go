package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, ans)
	*ans = append(*ans, root.Val)
	dfs(root.Right, ans)
}
