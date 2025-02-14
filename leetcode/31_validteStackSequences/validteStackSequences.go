package validteStackSequences

func validteStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed)+len(popped))
	i, j := 0, 0
	for ; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
