package main

func findDuplicate(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i]-1 != i {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return -1
}
