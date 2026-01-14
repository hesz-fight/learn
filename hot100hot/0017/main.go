package main

var dict = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCominations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans := []string{}
	backward(ans, []byte{}, 0, digits)
	return ans
}

func backward(ans []string, tmpC []byte, curInd int, digits string) {
	if curInd == len(digits) {
		ans = append(ans, string(tmpC))
		return
	}
	curByte := digits[curInd]
	for _, elem := range dict[curByte] {
		tmpC = append(tmpC, elem)
		backward(ans, tmpC, curInd+1, digits)
		tmpC = tmpC[:len(tmpC)-1]
	}
}
