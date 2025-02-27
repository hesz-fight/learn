package maxprofit

func maxProfit(prices []int) int {
	ans := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice < prices[i] {
			minPrice = prices[i]
			continue
		}
		if prices[i]-minPrice > ans {
			ans = prices[i] - minPrice
		}
	}

	return ans
}
