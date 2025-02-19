package findnthdigit

import "strconv"

// 一位数字：1-9；9个
// 二位数字：10-99；90个
// 三位数字：100-999；900个
// ...
// 思路：将不同位数分为不同的区间
// 1、先定位到所在区间，假设位数是x
// 2、用 n 减去 1~x 位数区间的数字和，假设结果是y
// 3、定位到x位数区间里的数字：10^X + (y - 1) / x
// 4、定位到数字的第几位
func findNthDigit(n int) int {
	x := 1
	startOfx := 1
	sumOfDigit := 0
	ans := 0
	for {
		// 计算x位数的数字和
		lastSumOfDigit := sumOfDigit
		sumOfDigit += startOfx * 9 * x
		if sumOfDigit >= n {
			// 找到了位数区间
			num := startOfx + (n-lastSumOfDigit-1)/x
			ind := (n - lastSumOfDigit - 1) % x
			numStr := strconv.Itoa(num)
			ans = int(numStr[ind]) - int('0')
			break
		}
		x++
		startOfx *= 10
	}
	return ans
}
