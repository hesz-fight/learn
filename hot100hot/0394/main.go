package main

import "strconv"

// 栈应用
// Last in First out
func decodeString(s string) string {
	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		if c != ']' {
			stack = append(stack, c)
			continue
		}

		end := len(stack) - 1
		for end >= 0 && stack[end] != '[' {
			end--
		}
		letter := append([]byte{}, stack[end+1:]...)
		// pop
		stack = stack[:end]

		end = len(stack) - 1
		for end >= 0 && isNumber(stack[end]) {
			end--
		}
		numb := append([]byte{}, stack[end+1:]...)
		stack = stack[:end+1]

		// push number * letter
		numInt, _ := strconv.Atoi(string(numb))
		for i := 0; i < numInt; i++ {
			stack = append(stack, letter...)
		}
	}

	return string(stack)
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}
