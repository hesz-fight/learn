package inordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 栈模拟
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	output := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}
		top := stack[len(stack)-1]
		output = append(output, top.Val)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
			root = top.Right
		}
	}
	return output
}
