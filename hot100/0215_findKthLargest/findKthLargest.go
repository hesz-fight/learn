package findkthlargest

import "math"

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), 0)
	}

	ans := 0
	for i := 0; i < k; i++ {
		ans = nums[0]
		nums[0] = math.MinInt
		heapify(nums, len(nums), 0)
	}

	return ans
}

func heapify(nums []int, size int, root int) {
	left := 2*root + 1
	right := left + 1
	maxValNode := root
	if left < size && nums[left] > nums[maxValNode] {
		maxValNode = left
	}
	if right < size && nums[right] > nums[maxValNode] {
		maxValNode = right
	}
	if maxValNode != root {
		nums[root], nums[maxValNode] = nums[maxValNode], nums[root]
		heapify(nums, size, maxValNode)
	}
}
