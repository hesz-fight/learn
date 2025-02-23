package iioccroncenum

func occrOnceNum(nums []int) int {
	// 统计每一个比特位1出现的次数
	bitSum := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if num&(1<<i) != 0 {
				bitSum[i]++
			}
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		bitSum[i] %= 3
		if bitSum[i] == 1 {
			ans = ans ^ 1<<i
		}

	}

	return ans
}
