package minpathsum

// dp建模
// 问题可以拆解为子问题、并且具有最优子结构
// dp定义：dp[i][j] 为到达位置{i,j}的最小路径和
// 递推公式： dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	dp := [][]int{}
	for i := 0; i < len(grid); i++ {
		dp = append(dp, make([]int, len(grid[i])))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
