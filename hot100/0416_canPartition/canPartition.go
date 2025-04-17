package canpartition

import "sort"

func canPartion(nums []int) bool {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	sumL, sumR := 0, 0
	for i <= j {
		if sumL <= sumR {
			sumL += nums[i]
			i++
		} else {
			sumR += nums[j]
			j--
		}
	}
	return sumL == sumR
}

// 转换成是否能凑成sum(nums)/2的0-1背包问题
func canPartion1(nums []int) bool {
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
