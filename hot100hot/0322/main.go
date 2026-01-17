package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// -1 代表无法兑换
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minVal := math.MaxInt
		for _, coin := range coins {
			if i < coin {
				continue
			}
			if dp[i-coin] != -1 {
				minVal = min(minVal, dp[i-coin]+1)
			}
		}
		if minVal != math.MaxInt {
			dp[i] = minVal
		}
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
