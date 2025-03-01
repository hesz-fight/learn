package strtoint

// strToInt ...
func strToInt(str string) int {
	chs := []byte(str)
	i := 0
	for i < len(chs) && chs[i] == ' ' {
		i++
	}
	if (chs[i] < '0' || chs[i] > '9') && chs[i] != '-' && chs[i] != '+' {
		return 0
	}
	sign := 1
	if chs[i] == '-' {
		sign = -1
		i++
	} else if chs[i] == '+' {
		i++
	}

	ans := 0
	for i < len(chs) && '0' <= chs[i] && chs[i] <= '9' {
		ans = ans*10 + int(chs[i]-'0')
		i++
	}

	return sign * ans
}
