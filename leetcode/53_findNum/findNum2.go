package findnum

// 找到左右边界
func findNum2(nums []int, target int) int {
	l1 := findRightBound(nums, target-1)
	l2 := findRightBound(nums, target)
	return l2 - l1
}

// 找到大于taget的第一个数
func findRightBound(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		k := i + (j-i)/2
		if nums[k] <= target {
			i = k + 1
		} else {
			// 目标是找到大于k的第一个数 收缩右边界不能-1 否则有可能漏掉答案
			j = k
		}
	}

	return j
}
