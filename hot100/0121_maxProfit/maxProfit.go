package maxprofit

func maxProfit(prices []int) int {
	ans := 0
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < -dp[i-1][0] {
			// 成本更低时才会买入
			dp[i][0] = -prices[i]
		} else {
			dp[i][0] = dp[i-1][0]
		}
		dp[i][1] = dp[i-1][0] + prices[i]
		ans = max(dp[i][1], ans)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
