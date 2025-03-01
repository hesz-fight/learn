package maxarea

// 头尾双指针
func maxArea(height []int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	// 头尾双指针
	ans, i, j := 0, 0, len(height)
	for i < j {
		if height[i] < height[j] {
			ans = max(height[i]*(j-i), ans)
			i++
		} else {
			ans = max(height[j]*(j-i), ans)
			j--
		}
	}
	return ans
}
