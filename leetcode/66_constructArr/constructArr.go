package constructarr

// 前缀积
func constructArr(a []int) []int {
	multiA := make([]int, len(a))
	multiB := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		j := len(a) - 1 - i
		if i == 0 {
			multiA[i] = 1
			multiB[j] = 1
			continue
		}
		multiA[i] = multiA[i-1] * a[i-1]
		multiB[j] = multiB[j+1] * a[j+1]
	}

	for i := 0; i < len(multiA); i++ {
		multiA[i] = multiA[i] * multiB[i]
	}

	return multiA
}
