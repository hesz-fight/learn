package main

import "sort"

// [-1,0,1,2,-1,-4]
// [-4,-1,-1,0,1,2]
func threeSum(nums []int) [][]int {
	// 排序去重复
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := [][]int{}
	n := len(nums)
	for i := 0; i <= n-3; i++ {
		j := i + 1
		k := n - 1
		if nums[k] > 0 {
			continue
		}
		// 跳过重复的 i
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// 跳过重复的 j 和 k
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}
