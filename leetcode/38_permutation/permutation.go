package permutation

import "strings"

var ans []string

// dps+回溯
func permutation(str string) []string {
	recur([]byte(str), []int{}, map[int]bool{})
	return ans
}

func recur(byteArr []byte, indArr []int, visited map[int]bool) {
	// base case
	if len(indArr) == len(byteArr) {
		sb := strings.Builder{}
		for _, i := range indArr {
			sb.WriteByte(byteArr[i])
		}
		ans = append(ans, sb.String())
	}
	// 去重复
	byteSet := make(map[byte]bool)
	for i := 0; i < len(byteArr); i++ {
		if byteSet[byteArr[i]] || visited[i] {
			continue
		}
		byteSet[byteArr[i]] = true
		// 更新条件
		visited[i] = true
		indArr = append(indArr, i)
		recur(byteArr, indArr, visited)
		// 条件回溯
		visited[i] = false
		indArr = indArr[:len(indArr)-1]
	}
}
