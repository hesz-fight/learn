package isinglenumbers

// 位运算：异或
// 0 ^ 0 = 0; 1 ^ 1 = 0; 0 ^ 1 = 1
func singleNumbers(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	// 出现两次的数比特位完全相同 已经被异或为 0
	// 因此 xorSum 是两个只出现了一次的数字的异或结果
	// 根据 xorSum 的第一个非零比特位 切分数组 使得两个不同的数落在不同的组里
	i := 0
	for (xorSum & (1 << i)) == 0 {
		i++
	}

	a, b := 0, 0
	for _, num := range nums {
		if (num & (1 << i)) == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
