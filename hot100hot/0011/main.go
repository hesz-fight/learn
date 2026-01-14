package main

// 头尾双指针
func maxArea(height []int) int {
	ans := 0
	head, tail := 0, len(height)-1
	for head < tail {
		if height[head] > height[tail] {
			ans = max(ans, height[tail]*(tail-head))
			tail--
		} else {
			ans = max(ans, height[head]*(tail-head))
			head++
		}
	}

	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
