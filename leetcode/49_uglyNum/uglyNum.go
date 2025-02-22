package uglynum

func uglyNum(n int) int {
	if n == 1 {
		return 1
	}
	uglyNumArr := make([]int, n+1)
	uglyNumArr[1] = 1
	recur(2, n, &uglyNumArr, 1, 1, 1)
	return uglyNumArr[n]
}

func recur(ind int, n int, uglyNumArr *[]int, ugly2 int, ugly3 int, ugly5 int) {
	if ind > n {
		return
	}
	num2 := (*uglyNumArr)[ugly2] * 2
	num3 := (*uglyNumArr)[ugly3] * 3
	num5 := (*uglyNumArr)[ugly5] * 5
	minUgly := num2
	if minUgly > num3 {
		minUgly = num3
	}
	if minUgly > num5 {
		minUgly = num5
	}
	(*uglyNumArr)[ind] = minUgly
	if minUgly == num2 {
		ugly2++
	}
	if minUgly == num3 {
		ugly3++
	}
	if minUgly == num5 {
		ugly5++
	}
	recur(ind+1, n, uglyNumArr, ugly2, ugly3, ugly5)
}
