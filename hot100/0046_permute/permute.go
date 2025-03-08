package permute

var (
	ans     [][]int
	tmp     []int
	visited map[int]struct{}
)

// 全排列：回溯算法
// 使用map统计，避免重复元素
func permute(nums []int) [][]int {
	{
		ans = [][]int{}
		tmp = []int{}
		visited = map[int]struct{}{}
	}
	dfs(nums)
	return ans
}

func dfs(nums []int) {
	if len(tmp) == len(nums) {
		ans = append(ans, append([]int{}, tmp...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		}

		tmp = append(tmp, nums[i])
		visited[nums[i]] = struct{}{}
		dfs(nums)
		delete(visited, nums[i])
		tmp = tmp[:len(tmp)-1]
	}
}
