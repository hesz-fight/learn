package main

func minDistance1(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化 某个字符串为空串时 将整个空串转换成另外一个字符串需要的操作数就等于另外这个字符串的长度
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			}
		}
	}

	// dp[i][j] 代表word1[:i]变成word1[:j]需要的最小操作
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// word1 删除末尾的元素使得长度减1
				// word1 新增一个和 word2 末尾相同的元素，使得 word2 长度减1
				// word1 修改末尾的元素使其等于 word2 末尾的元素，使得 word1 和 word2 的长度同时减1
				// 三种操作得到的子问题中的最小的解+1就是当前子问题的解
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
