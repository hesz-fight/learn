package canjump

var memo []int

func canJump(nums []int) bool {
	memo = make([]int, len(nums))
	return helper(nums, 0)
}

func helper(nums []int, ind int) bool {
	if ind >= len(nums)-1 {
		return true
	}
	if memo[ind] != 0 {
		return memo[ind] > 0
	}

	for i := 1; i <= nums[ind]; i++ {
		flag := helper(nums, ind+i)
		if flag {
			memo[ind+i] = 1
		} else {
			memo[ind+i] = -1
		}
		if flag {
			return true
		}
	}
	return false
}
