package maximalsquare

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] == 1 {
			ans = 1
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		if dp[0][j] == 1 {
			ans = 1
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				minVal := min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				dp[i][j] = minVal + 1
				ans = max(dp[i][j]*dp[i][j], ans)
			}
		}
	}

	return ans
}
