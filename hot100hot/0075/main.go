package main

func sortColors(nums []int) {
	zeroPos := 0
	twoPos := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[zeroPos], nums[i] = nums[i], nums[zeroPos]
			zeroPos++
		}
	}

	for i := len(nums) - 1; i >= zeroPos; i-- {
		if nums[i] == 2 {
			nums[twoPos], nums[i] = nums[i], nums[twoPos]
			twoPos--
		}
	}
}
