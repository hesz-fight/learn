package main

func canJump2(nums []int) bool {
	// dp[i] 代表从起点出发是否能到达i为终点
	dp := make([]bool, 0, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
