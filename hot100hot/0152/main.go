package main

func maxProduct(nums []int) int {
	n := len(nums)
	dp := [][2]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, [2]int{})
	}
	ans := nums[0]
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		dp[i][0] = min(dp[i-1][1]*x, min(dp[i-1][0]*x, x))
		dp[i][1] = max(dp[i-1][1]*x, max(dp[i-1][0]*x, x))
		if dp[i][1] > ans {
			ans = dp[i][1]
		}
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
