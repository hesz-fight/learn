package uglynum

// ugly = 2^i + 3^j + 5^k
// uglyNum2
func uglyNum2(n int) int {
	uglyNum := make([]int, n)
	uglyNum[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num2, num3, num5 := uglyNum[p2]*2, uglyNum[p3]*3, uglyNum[p5]*5
		minVal := minVal(num2, num3, num5)
		if minVal == num2 {
			p2++
		}
		if minVal == num3 {
			p3++
		}
		if minVal == num5 {
			p5++
		}
		uglyNum[i] = minVal
	}
	return uglyNum[n-1]
}

func minVal(x, y, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}

	return min
}
