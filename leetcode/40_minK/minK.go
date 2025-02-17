package mink

func minK(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return []int{}
	}
	l := 0
	r := len(input) - 1
	for l <= r {
		p := partion(input, l, r)
		if p+1 == k {
			return input[:p+1]
		} else if p+1 < k {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return []int{}
}

func minK2(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return []int{}
	}

	ans := []int{}
	// 1. 建立初始堆
	buildHeap(input)
	for i := 0; i < k; i++ {
		ans = append(ans, input[0])
		input[len(input)-1-i], input[0] = input[0], input[len(input)-1-i]
		heapifyPart(input, 0, len(input)-1-i)
	}

	return ans
}

// 冒泡排序
func bubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := len(input) - 1; j > i; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}
}

// 快速排序
func quickSort(input []int, l, r int) {
	if l >= r {
		return
	}
	j := partion(input, l, r)
	quickSort(input, l, j-1)
	quickSort(input, j+1, r)
}

// 快速排序：数组切分
func partion(input []int, l int, r int) int {
	pivot := input[r]
	j := l
	for i := l; i < r; i++ {
		if input[i] < pivot {
			input[i], input[j] = input[j], input[i]
			j++
		}
	}
	input[j], input[r] = input[r], input[j]
	return j
}

func heapSort(input []int) {
	// 1. 建立初始堆
	buildHeap(input)
	// 2. 后续不再需要从0-1建立堆 只需要交换后局部堆化就行
	for i := len(input) - 1; i > 0; i-- {
		input[i], input[0] = input[0], input[i]
		// 交换以后破坏了堆的结构 需要从最顶部的根节点往下做局部的堆化
		heapifyPart(input, 0, i)
	}
}

// 建立初始堆 树结构的冒泡
// 从最后一个非叶子节点开始 自底向上交换，直到将最小/最大值交换到堆顶部
func buildHeap(arr []int) {
	lastNotLeaf := len(arr)/2 - 1
	for i := lastNotLeaf; i >= 0; i-- {
		heapifyPart(arr, i, len(arr))
	}
}

// 将根节点、左节点、右节点的最大值交换到根节点的位置
func heapifyPart(arr []int, root int, heapSize int) {
	left := 2*root + 1
	right := left + 1
	min := root
	if left < heapSize && arr[left] < arr[min] {
		min = left
	}
	if right < heapSize && arr[right] < arr[min] {
		min = right
	}
	if min != root {
		arr[min], arr[root] = arr[root], arr[min]
		heapifyPart(arr, min, heapSize)
	}
}
