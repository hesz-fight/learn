package sortcolors

// 头尾双指针，将元素交换到合适的位置
func sortColors(nums []int) {
	p0 := 0
	p2 := len(nums) - 1
	// 通过i和p2控制循环
	for i := 0; i <= p2; i++ {
		// 交换直到nums[i]不为2
		for nums[i] == 2 && i <= p2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// 双指针
func sortColors2(nums []int) {
	p0 := 0
	p1 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p1 > p0 {
				// 此时 p0 和 p1 中间是连续的 1
				// 交换了 p0 和 i 会将一个 1 交换出去 需要再交换回来
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
