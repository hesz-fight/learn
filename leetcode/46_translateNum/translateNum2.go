package translatenum

import (
	"strconv"
)

// 动态规划
func translateNum2(num int) int {
	numByte := []byte(strconv.Itoa(num))
	dp := make([]int, len(numByte)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		// 翻译一个字符
		val := dp[i-1]
		// 翻译两个字符
		if numByte[i-2] == '1' || (numByte[i-2] == '2' && numByte[i-1] < '6') {
			val += dp[i-2]
		}
		dp[i] = val
	}
	return dp[len(dp)-1]
}
