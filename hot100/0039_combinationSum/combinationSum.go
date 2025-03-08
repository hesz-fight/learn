package combinationsum

var targetTmp int
var candidatesTmp []int
var ans [][]int

// 组合数：回溯
func combinationSum(candidates []int, target int) [][]int {
	candidatesTmp = candidates
	targetTmp = target
	traceback(target, []int{})
	return ans
}

func traceback(rem int, visited []int) {
	if rem < 0 {
		return
	}
	if rem == 0 {
		ans = append(ans, visited)
	}
	// rem > 0
	for _, can := range candidatesTmp {
		// 去重复 让 visited 里面的数字从小到大排列
		if len(visited) > 0 && can < visited[len(visited)-1] {
			continue
		}
		visited = append(visited, can)
		traceback(rem-can, visited)
		visited = visited[:len(visited)-1]
	}
}
