package main

// 求最大值、最小值
// 决策是动态做出的，每一步做一个局部最优的决策，最终得到全局最优的决策
// 问题是可以拆解的，具有最优子结构
func rob1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return maxFunc(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	// 初始化
	// 基本问题的解一般可以直接分析得到
	dp[0] = nums[0]
	dp[1] = maxFunc(nums[0], nums[1])
	// 遍历dp状态，对问题规模进行递推
	for i := 2; i < n; i++ {
		dp[i] = maxFunc(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
