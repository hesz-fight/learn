package longsubstring

// 最长不重复子字符串的长度
// 快慢双指针
func longsubstring(str string) int {
	slow, fast := 0, 0
	ans := 0
	visited := make(map[byte]bool)
	byteArr := []byte(str)
	for fast < len(byteArr) {
		if !visited[byteArr[fast]] {
			visited[byteArr[fast]] = true
			fast++
			continue
		}
		ans = max(ans, fast-slow)
		// 移动slow指针
		for !visited[byteArr[slow]] {
			slow++
		}
		slow++
		fast++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
