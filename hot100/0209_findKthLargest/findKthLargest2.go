package findkthlargest

func findKthLargest2(nums []int, k int) int {
	return quickSort(nums, k, 0, len(nums)-1)
}

func quickSort(nums []int, k int, i int, j int) int {
	if i == j {
		return nums[k]
	}
	mid := divide2(nums, i, j)
	// 第k大的数下标是k-1
	if mid == k-1 {
		return nums[mid]
	} else if mid < k-1 {
		return quickSort(nums, k, mid+1, j)
	} else {
		return quickSort(nums, k, i, mid-1)
	}
}

// 极端case过不了
func divide(nums []int, i int, j int) int {
	p := nums[j]
	slow := i
	for fast := i; fast <= j-1; fast++ {
		if nums[fast] > p {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
	nums[slow], nums[j] = nums[j], nums[slow]
	return slow
}

func divide2(nums []int, i int, j int) int {
	p := nums[i]
	l := i + 1
	r := j
	for l <= r {
		for l < len(nums) && nums[l] > p {
			l++
		}
		for r >= 0 && nums[r] < p {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	nums[i], nums[r] = nums[r], nums[i]

	return r
}
