package iimissingnum

func fineMissingNum(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := i + (j-1)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			// -1 会漏掉答案
			j = mid
		}
	}
	return j
}
