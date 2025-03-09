package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
		for i := 0; i < size; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		ans = append(ans, tmp)
	}
	return ans
}
