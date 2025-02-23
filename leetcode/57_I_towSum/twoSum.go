package towsum

func twoSum(nums []int, tar int) []int {
	head := 0
	tail := len(nums) - 1
	for head < tail {
		if nums[head]+nums[tail] < tar {
			head++
			continue
		}
		if nums[head]+nums[tail] > tar {
			tail--
			continue
		}
		return []int{nums[head], nums[tail]}
	}

	return []int{}
}
