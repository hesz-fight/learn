package verifypostorder

func verifyPostorder(in []int) bool {
	return recur(in, 0, len(in)-1)
}

// 分支递归 递归区间
func recur(in []int, i int, j int) bool {
	if i >= j {
		return true
	}

	// 左子树和右子树的边界
	p := i
	for in[i] < in[j] {
		p++
	}

	//
	q := p
	for in[q] > in[j] {
		q++
	}

	return q == j && recur(in, i, p-1) && recur(in, p, j-1)
}
