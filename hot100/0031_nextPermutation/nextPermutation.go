package nextpermutation

// 从右往前找一个较小数：找到第一个相邻的顺序对
// 在较小数的右区间找到一个较大数 交换两个数
// 重排较小数右区间 得到下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	// 从右往前找到第一个相邻的顺序对nums[i] nums[j]
	i := n - 2
	for i > 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 在较小数的右区间找到一个较大数
	j := n - 1
	for j > 0 && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	// 重排i的右区间
	for p, q := i+1, n-1; p < q; p, q = p+1, q-1 {
		nums[p], nums[q] = nums[q], nums[p]
	}
}
