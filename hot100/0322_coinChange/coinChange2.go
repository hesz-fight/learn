package coinchange

import "math"

func coinChange2(coins []int, amount int) int {
	minFunc := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		cnt := math.MaxInt
		for _, coin := range coins {
			if coin <= amount && dp[amount-coin] != -1 {
				cnt = minFunc(dp[amount-coin]+1, cnt)
			}
		}
		if cnt != math.MaxInt {
			dp[i] = cnt
			continue
		}
		// 凑不出当前面额
		dp[i] = -1
	}

	return dp[amount]
}
