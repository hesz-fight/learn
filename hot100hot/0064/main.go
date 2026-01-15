package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j > 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else if j == 0 {
				if i > 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
