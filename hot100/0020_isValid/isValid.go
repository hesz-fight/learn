package isvalid

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		topElem := stack[len(stack)-1]
		if (topElem == '(' && v == ')') ||
			(topElem == '[' && v == ']') ||
			(topElem == '{' && v == '}') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
