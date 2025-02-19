package maxsubarray

import "math"

func maxsubarray(nums []int) int {
	maxVal := math.MinInt
	for l := 1; l <= len(nums); l++ {
		for i := 0; i+l-1 < len(nums); i++ {
			s := 0
			for j := i; j < l+i; j++ {
				s += nums[j]
			}
			if s > maxVal {
				maxVal = s
			}
		}
	}
	return maxVal
}

func maxsubarray2(nums []int) int {
	dp := make([]int, 0, len(nums))
	ans := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
