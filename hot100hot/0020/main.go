package main

// "()[]{}"
func isValid(s string) bool {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st = append(st, s[i])
		} else if s[i] == ')' && st[len(st)-1] == '(' {
			st = st[:len(st)-1]
		} else if s[i] == ']' && st[len(st)-1] == '[' {
			st = st[:len(st)-1]
		} else if s[i] == '}' && st[len(st)-1] == '{' {
			st = st[:len(st)-1]
		}
	}

	return len(st) == 0
}
