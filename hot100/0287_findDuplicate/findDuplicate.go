package findduplicate

func findDuplicate(nums []int) int {
	// nums[i] == i + 1
	var ans int
Loop:
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}
		for nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				ans = nums[i]
				break Loop
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return ans
}
