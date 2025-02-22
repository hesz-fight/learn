package translatenum

import "strconv"

// 记忆化递归
var meno []int

func translateNum(num int) int {
	meno = make([]int, num)
	return recur([]byte(strconv.Itoa(num)), 0)
}

// 当前开始的下标
func recur(nums []byte, curInd int) int {
	// base case
	if curInd >= len(nums) {
		return 1
	}
	if meno[curInd] != 0 {
		return meno[curInd]
	}
	// recur case
	rtn := recur(nums, curInd+1)
	if curInd+1 < len(nums) &&
		(nums[curInd] == '1' || (nums[curInd] == '2' && nums[curInd+1] <= '5')) {
		rtn += recur(nums, curInd+2)
	}
	meno[curInd] = rtn

	return meno[curInd]
}
