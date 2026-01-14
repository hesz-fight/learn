package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	// 双指针合并区间
	i := 0
	j := 0
	ans := [][]int{}

	end := intervals[0][1]
	for j < len(intervals) {
		if intervals[j][0] <= intervals[i][1] {
			j++
			end = max(end, intervals[j][1])
		} else {
			// i 到 j - 1 可以合并成一个区间
			ans = append(ans, []int{intervals[i][0], intervals[j-1][1]})
			// 更新i为j
			i = j
			// 更新端点值
			end = intervals[0][i]
		}
	}
	ans = append(ans, []int{intervals[i][0], intervals[j-1][1]})

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
