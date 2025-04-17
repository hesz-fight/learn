package coinchange

import "math"

var memo map[int]int

/**
 * 备忘录递归。这种递归方法只能只能在回退阶段做值记录。
 */
func coinChange(coins []int, amount int) int {
	return helper(coins, amount)
}

// 暴力递归。非尾递归形式。回退阶段不做值计算的递归形式叫做尾递归
func helper(coins []int, sum int) int {
	if sum == 0 {
		return 0
	}
	if _, ok := memo[sum]; ok {
		return memo[sum]
	}
	cnt := math.MaxInt
	for _, coin := range coins {
		if coin > sum {
			continue
		}
		tmpCnt := helper(coins, sum-coin)
		if tmpCnt < cnt {
			cnt = tmpCnt
		}
	}
	memo[sum] = -1

	if cnt != math.MaxInt {
		memo[sum] = cnt + 1
	}

	return memo[sum]
}
