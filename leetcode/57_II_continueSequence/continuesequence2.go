package iicontinuesequence

func continuesequence2(target int) [][]int {
	ans := [][]int{}
	sum := make([]int, target+1/2)
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + i
	}
	dict := map[int]int{}
	for t, s := range sum {
		dict[s] = t
		if s-target <= 0 {
			continue
		}
		if _, ok := dict[s-target]; ok {
			ans = append(ans, genArr(dict[s-target]+1, t))
		}
	}

	return ans
}
