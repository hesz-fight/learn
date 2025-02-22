package longsubstring

func longsubstring2(str string) int {
	ans := 0
	charArr := []byte(str)
	char2ind := make(map[byte]int)
	dp := make([]int, len(charArr))
	dp[0] = 1
	for j := 1; j < len(charArr); j++ {
		i := char2ind[charArr[j]]
		// 	j-i 是当前字符加入最长不重复子串的解
		if j-i > dp[j-1] {
			dp[i] = j - i
		} else {
			// 	当前字符不加入 即当前字符已经重复出现的情况
			dp[i] = dp[j-1] + 1
		}
		char2ind[charArr[j]] = j
		ans = max(ans, dp[i])
	}

	return ans
}
