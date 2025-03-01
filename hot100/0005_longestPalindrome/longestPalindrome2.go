package longestpalindrome

// dp 区间动态规划
func longestPalindrome2(s string) string {
	n := len(s)
	// s[i:j]回文串的长度
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}
	maxLen, l, r := 1, 0, 0
	for m := 1; m <= n; m++ { // 区间长度
		for i := 0; i <= n-m; i++ { // 左边界
			j := i + m - 1
			if s[i] != s[j] {
				dp[i][j] = 1
			} else {
				if m <= 3 {
					dp[i][j] = m

				} else {
					// s[i+1:j-1]不是回文串
					if dp[i+1][j-1] == 1 {
						dp[i+1][j-1] = 1
					} else {
						dp[i][j] = dp[i+1][j-1] + 2
					}
				}
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				l, r = i, j
			}
		}
	}
	return s[l : r+1]
}
