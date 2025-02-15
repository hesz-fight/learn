package ilevelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 按层打印二叉树
func ilevelorder(root *TreeNode) []int {
	ans := make([]int, 0, 64)
	queue := make([]*TreeNode, 0, 64)
	queue = append(queue, root)
	// 队列长度动态变化
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		ans = append(ans, curNode.Val)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}

	return ans
}
