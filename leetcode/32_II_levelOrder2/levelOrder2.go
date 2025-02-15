package iilevelorder2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ilevelorder2(root *TreeNode) [][]int {
	ans := make([][]int, 0, 64)
	queue := make([]*TreeNode, 0, 64)
	queue = append(queue, root)
	for len(queue) > 0 {
		len := len(queue)
		tmp := make([]int, 0, len)
		for i := 0; i < len; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, tmp)
	}

	return ans
}
