package threesum

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 排序是为了去重
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := [][]int{}
	// 固定第一个数 i
	// 用头尾双指针去遍历第二和第三个数 j 和 k
	for i := 0; i < len(nums); i++ {
		// 元素不足 3 个 时 提前结束
		if i > len(nums)-3 {
			break
		}
		// 大于 0 时跳过 因为 nums[k] >= num[j] >= nums[i]
		// 当 nums[i] > 0 时, nums[k] + num[j] + nums[i] > 0
		if nums[i] > 0 {
			break
		}
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 在一个有序数组里求两数之和
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
			// 根据值移动指针 不能跳过任何可能结果
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
				k--
				// 跳过重复元素
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// 跳过重复元素
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}
