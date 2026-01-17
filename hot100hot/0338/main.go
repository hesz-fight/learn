package main

func countBits(n int) []int {
	rtn := make([]int, 0, n+1)
	for i := 0; i < n+1; i++ {
		rtn = append(rtn, helper(i))
	}
	return rtn
}

func helper(num int) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
