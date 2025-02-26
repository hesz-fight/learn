package irotateword

func rotateWord(words []string) []string {
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return words
}
