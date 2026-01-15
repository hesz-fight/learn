package main

var memo = map[string]int{}

// 自顶向下拆解子问题：递归
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		// 一旦某个单词为空串
		// 另外一个单词变成另外一个单词的操作数就是非空单词的长度
		return max(len(word1), len(word2))
	}

	key := word1 + ":" + word2
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	// 如果最后一个单词相等 不用做任何操作 可以同时消掉当前单词
	if word1[len(word1)-1] == word2[len(word2)-1] {
		return minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
	}
	// 插入、删除、修改之中的最小值
	v1 := minDistance(word1, word2[:len(word2)-1])
	v2 := minDistance(word1[:len(word1)-1], word2)
	v3 := minDistance(word1[:len(word1)-1], word2[:len(word2)-1])

	memo[key] = min(v1, min(v2, v3)) + 1

	return memo[key]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
