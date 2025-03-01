package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	uniqWin := make(map[byte]int, len(s))
	slow, fast, ans := 0, 0, 0
	for fast < len(s) {
		ind, ok := uniqWin[s[fast]]
		if ok {
			// 移动左边界 防止往回跑需要求一个最大值
			slow = max(slow, ind+1)
		}
		uniqWin[s[fast]] = fast
		ans = max(fast-slow+1, ans)
		fast++
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
