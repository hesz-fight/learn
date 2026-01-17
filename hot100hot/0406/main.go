package main

// 多维度排序问题
// 对第一个维度降序排序
// 在第一个维度相同的时候，对第二个维度升序排序
import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0]-people[j][0] > 0
		} else {
			return people[i][1]-people[j][1] < 0
		}
	})
	rtn := [][]int{}
	for _, p := range people {
		rtn = append(rtn[:p[1]], append([][]int{p}, rtn[p[1]:]...)...)
	}
	return rtn
}
