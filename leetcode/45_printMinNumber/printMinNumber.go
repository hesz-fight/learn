package printminnumber

import (
	"strconv"
	"strings"
)

func PrintMinNumber(numbers []int) string {
	quickSort(numbers, 0, len(numbers)-1)
	sb := strings.Builder{}
	for _, num := range numbers {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func quickSort(numbers []int, l, r int) {
	if l >= r {
		return
	}
	p := partion(numbers, l, r)
	quickSort(numbers, l, p-1)
	quickSort(numbers, p+1, r)
}

func partion(numbers []int, l, r int) int {
	pv := numbers[r]
	slow, fast := l, l
	for fast < r {
		if less(numbers[fast], pv) {
			numbers[slow], numbers[fast] = numbers[fast], numbers[slow]
			slow++
		}
		fast++
	}
	numbers[slow], numbers[r] = numbers[r], numbers[slow]
	return slow
}

// 11 2
// 112 < 211
func less(num1 int, num2 int) bool {
	numStr1 := strconv.Itoa(num1)
	numStr2 := strconv.Itoa(num2)
	sum1, _ := strconv.Atoi(numStr1 + numStr2)
	sum2, _ := strconv.Atoi(numStr2 + numStr1)
	return sum1 < sum2
}
