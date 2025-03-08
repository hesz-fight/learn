package searchrange

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	i := binarySearch(nums, target-1)
	j := binarySearch(nums, target)
	if i == j {
		return []int{-1, -1}
	} else {
		return []int{i, j - 1}
	}
}

// 二分查找：大于targer的第一个数
func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	// 找边界时不等于
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] > target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if nums[i] == target {
		i++
	}

	return i
}
