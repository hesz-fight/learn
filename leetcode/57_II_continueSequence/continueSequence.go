package iicontinuesequence

// 滑动窗口
func continuesequence(target int) [][]int {
	ans := [][]int{}
	sum := 0
	l := 1
	for r := 2; r <= (target+1)/2; r++ {
		sum += r
		if sum < target {
			continue
		}
		// sum >= target
		for sum >= target {
			if sum == target {
				ans = append(ans, genArr(l, r))
				continue
			}
			sum -= l
			l++
		}
	}

	return ans
}

func genArr(l, r int) []int {
	arr := []int{}
	for i := l; i <= r; i++ {
		arr = append(arr, i)
	}
	return arr
}
