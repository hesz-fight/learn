package main

import "fmt"

/*
最长公共子串
题目要求：
给定两个字符串，找出它们的最长公共子串，并返回该子串的长度。

示例1：

输入：s1 = "abcdef", s2 = "zbcdf"

输出：3

解释：最长公共子串是 "bcd"，长度为 3。

示例2：

输入：s1 = "abc", s2 = "def"

输出：0

解释：没有公共子串。
*/

func main() {
	fmt.Println(maxCommonSubstring("abcdef", "zbcdf"))
	fmt.Println(maxCommonSubstring("abc", "def"))
}

func maxCommonSubstring(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
