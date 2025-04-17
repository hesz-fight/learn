package minwindow

func minWindow(s string, t string) string {
	window := map[byte]int{}
	mm := map[byte]int{}
	for i := 0; i < len(t); i++ {
		mm[t[i]]++
	}
	ans := ""
	l, r := 0, 0
	for l < len(s) && r < len(s) {
		if !isEqual(window, mm) {
			if _, ok := mm[s[r]]; ok {
				window[s[r]]++
			}
			r++
		} else {
			if ans == "" || r-l < len(ans) {
				ans = s[l:r]
			}
			if _, ok := mm[s[l]]; ok {
				window[s[l]]--
			}
			l++
		}
	}

	if isEqual(window, mm) && r-l < len(ans) {
		ans = s[l:r]
	}

	return ans
}

func isEqual(m1 map[byte]int, m2 map[byte]int) bool {
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			return false
		}
		if m1[k] < v {
			return false
		}
	}

	return true
}
