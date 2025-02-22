package longsubstring

// 使用map避免二次遍历
func longsubstring1(str string) int {
	slow, fast := 0, 0
	ans := 0
	// byte to index
	visited := make(map[byte]int)
	byteArr := []byte(str)
	for fast < len(byteArr) {
		cur := byteArr[fast]
		if _, ok := visited[cur]; ok {
			// 指向重复元素的下一个坐标
			// 防止左边界往回跑 求一个最大值
			slow = max(slow, visited[cur]+1)
		}
		ans = max(ans, fast-slow+1)
		// 更新字符最后一次出现的位置
		visited[cur] = fast
		fast++
	}

	return ans
}
