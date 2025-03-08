package merge

import "sort"

// 排序+双指针
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0, len(intervals))
	// 双层排序 将小区间放前面
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0]-intervals[j][0] < 0
		} else {
			return intervals[i][1]-intervals[j][1] < 0
		}
	})
	// 双指针遍历
	i, j := 0, 0
	for i < len(intervals) {
		// 新区间的右端点
		end := intervals[i][1]
		// 从下一个区间开始遍历 更新右端点
		j = i + 1
		// 判断下一个区间与上一个区间有没有重叠
		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		ans = append(ans, []int{intervals[i][0], end})
		// 更新下一个起点区间
		i = j
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
