package lettercombinations

var ans []string
var digit2letter map[byte][]byte

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans = []string{}
	digit2letter = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	backward(0, []byte(digits), make([]byte, 0, len(digits)))

	return ans
}

func backward(start int, digits []byte, rst []byte) {
	if start == len(digits) {
		ans = append(ans, string(rst))
		return
	}

	letter := digit2letter[digits[start]]
	for _, l := range letter {
		rst = append(rst, l)
		backward(start+1, digits, rst)
		rst = rst[:len(rst)-1]
	}
}
