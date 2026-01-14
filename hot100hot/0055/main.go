package main

func canJump(nums []int) bool {
	maxInd := 0
	for i, num := range nums {
		if i <= maxInd {
			maxInd = max(maxInd, i+num)
		}
		if maxInd >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
