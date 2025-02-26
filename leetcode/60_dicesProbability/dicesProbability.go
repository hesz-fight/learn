package dicesprobability

func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ {
		next := make([]float64, 5*n+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				next[j+k] += dp[j] / 6.0
			}
		}
		dp = next
	}

	return dp
}
