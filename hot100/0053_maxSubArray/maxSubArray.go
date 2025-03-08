package maxsubarray

import "math"

func maxSubarray(nums []int) int {
	ans := math.MinInt
	dp := make([]int, 0, len(nums))
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
