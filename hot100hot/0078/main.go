package main

func subsets(nums []int) [][]int {
	ans := [][]int{}
	tracebaack(nums, []int{}, &ans, 0)
	return ans
}

func tracebaack(nums []int, tmp []int, ans *[][]int, pos int) {
	*ans = append(*ans, append([]int{}, tmp...))
	for i := pos; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		tracebaack(nums, tmp, ans, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
