package mindistance

var memo = map[string]int{}

// 自顶向下拆解子问题：递归
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		// 一旦某个单词为空串
		// 另外一个单词变成另外一个单词的操作数就是非空单词的长度
		return max(len(word1), len(word2))
	}

	key := word1 + ":" + word2
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	// 如果最后一个单词相等 不用做任何操作 可以同时消掉当前单词
	if word1[len(word1)-1] == word2[len(word2)-1] {
		return minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
	}
	// 插入、删除、修改之中的最小值
	v1 := minDistance(word1, word2[:len(word2)-1])
	v2 := minDistance(word1[:len(word1)-1], word2)
	v3 := minDistance(word1[:len(word1)-1], word2[:len(word2)-1])

	memo[key] = min(v1, min(v2, v3)) + 1

	return memo[key]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 自底向上递推子问题：动态规划
// dp[i][j]: words1[0:i]变成words2[0:j]需要的最少操作
// 递推公式：
// if words1[i] == words2[j] dp[i][j] = dp[i-1][j-1]
// else dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1])) + 1
func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], min(dp[i][j+1], dp[i][j])) + 1
			}
		}
	}
	return dp[m][n]
}
