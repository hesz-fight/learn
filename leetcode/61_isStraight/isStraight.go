package isstraight

import "math"

func isStraight(nums []int) bool {
	min, max := math.MaxInt, math.MinInt
	mp := map[int]bool{}
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if mp[num] {
			return false
		}
		mp[num] = true

		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max-min < 5
}
