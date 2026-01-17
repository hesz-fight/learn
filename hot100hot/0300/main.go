package main

import "math"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	ans := -1
	dp[0] = 1
	for i := 1; i < n; i++ {
		tmp := math.MaxInt
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				tmp = max(tmp, dp[j]+1)
			}
		}
		if tmp != math.MaxInt {
			dp[i] = tmp
		} else {
			dp[i] = 1
		}
		ans = max(ans, dp[i])
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return x
	}

	return y
}
