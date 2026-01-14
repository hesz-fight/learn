package main

func towSum(nums []int, target int) []int {
	set := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := set[target-nums[i]]; ok {
			return []int{set[set[target-nums[i]]], i}
		}
		set[nums[i]] = i
	}

	return []int{}
}
