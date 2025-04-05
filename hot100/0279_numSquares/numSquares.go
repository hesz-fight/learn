package numsquares

import "math"

// 动态规划
func numSquares(n int) int {
	var minFunc func(x, y int) int
	minFunc = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		minVal := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minVal = minFunc(minVal, dp[i-j*j]+1)
		}
	}
	return dp[n]
}
