package canjump

// 动态规划
func canJump2(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			break
		}
		for j := 1; j <= nums[i] && i+j <= len(nums)-1; j++ {
			dp[i+j] = true
		}
	}

	return dp[len(dp)-1]
}

// 贪心算法：在每一个位置尽可能跳到最远 更新最远可到达的距离
func canJump3(nums []int) bool {
	reach := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > reach {
			reach = i + nums[i]
			if reach >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
