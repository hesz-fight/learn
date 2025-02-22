package firstuniqchar

func uniqueChar(str string) byte {
	dict := make([]int, 25)
	for _, ch := range []byte(str) {
		dict[ch-'a'] += 1
	}
	for _, ch := range []byte(str) {
		if dict[ch-'a'] == 1 {
			return ch
		}
	}

	return ' '
}
