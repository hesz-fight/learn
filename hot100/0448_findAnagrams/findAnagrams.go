package finddisappearednumbers

// 滑动窗口
// map 统计数量+快速判断是否在窗口里面
func findAnagrams(s string, p string) []int {
	dict := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		dict[s[i]]++
	}
	ans := make([]int, 0)
	win := make(map[byte]int)
	i, j := 0, 0
	for j < len(s) {
		win[s[j]]++
		if j >= len(p) {
			win[s[i]]--
			i++

		}
		j++
		if isEqualMap(win, dict) {
			ans = append(ans, i)
		}
	}

	return ans
}

func isEqualMap(a, b map[byte]int) bool {
	for i := 0; i < 26; i++ {
		if a[byte('a'+i)] != b[byte('a'+i)] {
			return false
		}
	}
	return true
}
