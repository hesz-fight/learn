package main

var memo []int

func rob(nums []int) int {
	memo = make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	return recur(len(nums)-1, nums)
}

func recur(i int, nums []int) int {
	if i == 0 {
		return nums[0]
	} else if i == 1 {
		return max(nums[0], nums[1])
	}

	if memo[i] < 0 {
		memo[i] = max(recur(i-1, nums), recur(i-2, nums)+nums[i])
	}

	return memo[i]
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
