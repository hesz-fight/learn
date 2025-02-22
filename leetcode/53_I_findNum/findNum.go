package findnum

func findNum(nums []int, target int) int {
	k := binarySearch(nums, target)
	i, j := k-1, k+1
	ans := 1
	for i >= 0 {
		if nums[i] != target {
			break
		}
		ans++
		i--
	}
	for j < len(nums) {
		if nums[j] != target {
			break
		}
		ans++
		j--
	}

	return ans
}

func binarySearch(nums []int, tar int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		k := i + (j-i)/2
		if nums[k] == tar {
			return k
		}
		// bound
		if nums[k] < tar {
			i = k + 1
		} else {
			j = k - 1
		}
	}

	return -1
}
