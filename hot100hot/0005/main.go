package main

// babad
// babbad
func longestPalindrome(s string) string {
	maxLen := 0
	l, r := 0, 0
	n := len(s)
	dp := [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, []int{})
	}
	// dp[i][j] --> s[i:j] 以 i、j 为端点的子串的最长回文子串

	// 子串长度
	for len := 1; len <= n; len++ {
		// 左端点
		for i := 0; i <= n-len; i++ {
			// 右端点
			j := i + len - 1
			if s[i] != s[j] {
				dp[i][j] = 1
			} else {
				// s[i] == s[j]
				if len == 1 || len == 2 || len == 3 {
					dp[i][j] = len
				} else {
					// 如果减去当前端点剩余的字符串不是回文子串 则当前子串s[i:j]整体也不是回文子串
					if dp[i+1][j-1] == 1 {
						dp[i][j] = 1
					} else {
						dp[i][j] = dp[i+1][j-1] + 2
					}
				}
			}

			// 计算最大长度以及保存 i、j 的位置
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				l = i
				r = j
			}
		}
	}
	return s[l : r+1]
}
