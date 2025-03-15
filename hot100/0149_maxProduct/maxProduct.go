package maxproduct

import "math"

func maxProduct(nums []int) int {
	ans := math.MinInt
	n := len(nums)
	fMax := make([]int, n)
	fMin := make([]int, n)
	fMax[0], fMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		// 把 x 加到右端点为 i-1 的（乘积最大/最小）子数组后面，
		// 或者单独组成一个子数组，只有 x 一个元素
		fMax[i] = max(fMax[i-1]*x, max(fMin[i-1]*x, x))
		fMin[i] = min(fMax[i-1]*x, min(fMin[i-1]*x, x))
		if fMax[i] > ans {
			ans = fMax[i]
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
