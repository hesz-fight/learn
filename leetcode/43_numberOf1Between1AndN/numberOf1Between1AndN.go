package numberof1between1andn

// o(n*logn)
func numberOf1Between1AndN(n int) int {
	count := 0
	for m := 1; m <= n; m++ {
		i := m
		for i != 0 {
			if i%10 == 1 {
				count++
			}
			i /= 10
		}
	}
	return count
}

// o(n*logn)
func numberOf1Between1AndN2(n int) int {
	count := 0
	weight := 1
	for n != 0 {
		round := n / 10
		remain := n % 10
		count += round * weight
		if remain > 1 {
			count += weight
		} else if remain == 1 {
			count += n%weight + 1
		}
		n /= 10
		weight *= 10
	}

	return count
}
