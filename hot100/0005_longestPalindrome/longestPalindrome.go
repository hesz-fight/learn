package longestpalindrome

// 中心扩散 遍历中心点
func longestPalindrome(s string) string {
	l, r := 0, 0
	ans := 1
	for i := 0; i < len(s); i++ {
		// 中心点为i
		j, k := i, i
		cur := k - j + 1
		for j-1 >= 0 && k+1 <= len(s)-1 && s[j-1] == s[k+1] {
			j--
			k++
			cur += 2
		}
		if cur > ans {
			ans = cur
			l = j
			r = k
		}
		// 中心点为i和i+1 可预先判断减少遍历
		j, k = i, i
		if k+1 < len(s) && s[k+1] == s[i] {
			k++
			cur = k - j + 1
			for j-1 >= 0 && k+1 <= len(s)-1 && s[j-1] == s[k+1] {
				j--
				k++
				cur += 2
			}
			if cur > ans {
				ans = cur
				l = j
				r = k
			}
		}
	}

	return s[l : r+1]
}
