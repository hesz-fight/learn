package main

func subsets1(nums []int) [][]int {
	queue := [][]int{}
	// 放入初始的空集
	queue = append(queue, []int{})
	// 遍历每一个数放入队列的所有的集合中
	for _, num := range nums {
		// 当前队列的长度
		n := len(queue)
		for j := 0; j < n; j++ {
			tmp := append([]int{}, queue[j]...)
			tmp = append(tmp, num)
			queue = append(queue, tmp)
		}
	}
	return queue
}
