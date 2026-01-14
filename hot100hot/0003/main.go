package main

// abcad
func lengthOfLongestSubstring(s string) int {
	byte2ind := map[byte]int{}
	ans := 0
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		ind, ok := byte2ind[s[fast]]
		if ok {
			slow = max(slow, ind+1)
		}
		byte2ind[s[fast]] = fast
		ans = max(ans, fast-slow+1)
	}

	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
