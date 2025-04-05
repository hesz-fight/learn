package lengthoflis

func lengthOfLIS(nums []int) int {
	var maxFunc func(x, y int) int
	maxFunc = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		maxVal := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxVal = maxFunc(dp[j]+1, maxVal)
			}
		}
		dp[i] = maxVal
	}

	return dp[len(nums)-1]
}
