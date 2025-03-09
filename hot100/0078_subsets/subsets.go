package subsets

// BFS
func subsets(nums []int) [][]int {
	queue := make([][]int, 0, 16)
	queue = append(queue, []int{})
	for i := 0; i < len(nums); i++ {
		n := len(queue)
		for j := 0; j < n; j++ {
			tmp := append([]int{}, queue[j]...)
			tmp = append(tmp, nums[i])
			queue = append(queue, tmp)
		}
	}
	return queue
}
