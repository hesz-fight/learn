package add

// 无进位和 a^b
// 有进位和 a&b << 1
func add(a int, b int) int {
	if a == 0 || b == 0 {
		return a ^ b
	}
	// 有进位和最终会变成0
	return add(a^b, (a&b)<<1)
}
