package main

// 转换成是否能凑成sum(nums)/2的0-1背包问题
func canPartion(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 1; i <= sum; i++ {
		for _, num := range nums {
			if i < num {
				continue
			}
			if dp[i-num] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sum]
}
