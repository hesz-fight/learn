package main

type FreqNum struct {
	Freq int
	Num  int
}

func topKFrequent(nums []int, k int) []int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}

	freqNums := []*FreqNum{}
	for num, freq := range dict {
		freqNums = append(freqNums, &FreqNum{
			Freq: freq,
			Num:  num,
		})
	}

	// 从最后一个非叶子节点往前做堆化
	for i := k/2 - 1; i >= 0; i-- {
		heapfy(i, k, freqNums)
	}

	for i := k; i < len(freqNums); i++ {
		if freqNums[i].Freq > freqNums[0].Freq {
			freqNums[0] = freqNums[i]
			heapfy(0, k, freqNums)
		}
	}

	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, freqNums[i].Num)
	}

	return ans
}

func heapfy(root int, size int, freqNums []*FreqNum) {
	left := 2*root + 1
	right := left + 1
	minRoot := root
	if left < size && freqNums[left].Freq < freqNums[minRoot].Freq {
		minRoot = left
	}
	if right < size && freqNums[right].Freq < freqNums[minRoot].Freq {
		minRoot = right
	}
	if minRoot != root {
		freqNums[minRoot], freqNums[root] = freqNums[root], freqNums[minRoot]
		heapfy(minRoot, size, freqNums)
	}
}
