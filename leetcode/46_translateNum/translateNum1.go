package translatenum

import "strconv"

// 暴力枚举
func translateNum1(num int) int {
	return recur([]byte(strconv.Itoa(num)), 0)
}

// 当前开始的下标
func recur1(nums []byte, curInd int) int {
	// base case
	if curInd >= len(nums) {
		return 1
	}
	// recur case
	rtn := recur(nums, curInd+1)
	if curInd+1 < len(nums) &&
		(nums[curInd] == '1' || (nums[curInd] == '2' && nums[curInd+1] <= '5')) {
		rtn += recur(nums, curInd+2)
	}

	return rtn
}
