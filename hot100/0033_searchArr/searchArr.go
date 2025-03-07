package searcharr

// [4,5,6,7,0,1,2] 0
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		// 中点值和右端点比较，看中点落在哪个区间
		if nums[mid] < nums[j] {
			if nums[mid] > target {
				j = mid - 1
			} else {
				if nums[j] == target {
					return j
				} else if target < nums[j] {
					i = mid + 1
				} else {
					j = mid - 1
				}
			}
		} else {
			// nums[mid] > nums[j]
			if nums[mid] < target {
				i = mid + 1
			} else {
				if nums[j] == target {
					return j
				} else if target > nums[j] {
					j = mid - 1
				} else {
					i = mid + 1
				}
			}
		}
	}

	return -1
}
