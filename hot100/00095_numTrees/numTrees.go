package numtrees

// 此题难点在于找出问题的规律
// 当根节点为i时，左子树有[0,...,i]个节点，右子树有[i+1,...,n]个节点
// 则有：以i为根节点的二叉搜索树的个数为
// 左子树可以组成的二叉搜索树的个数乘上右子树可以组成的二叉搜索树的个数
var memo []int

func numTrees(n int) int {
	memo = make([]int, n+1)
	return recur(1, n)
}

// 返回节点[left, right]的二叉搜索树的个数
func recur(left, right int) int {
	if left >= right {
		return 1
	}
	if memo[right-left+1] != 0 {
		return memo[right-left+1]
	}
	rtn := 0
	for i := left; i <= right; i++ {
		l := recur(left, i-1)
		r := recur(i+1, right)
		rtn += l * r
	}
	memo[right-left+1] = rtn
	return rtn
}

// 动态规划
// 1. 问题是否能拆解？
// 2. 是否具有最优子结构？
// 生成dp数组:dp[问题的规模] = 问题的解
// 初始化dp数组：找出最小规模问题的解，一般是很直观的
// 增大问题的规模：解出每一个子问题的解
func numTrees2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// j是左子树的节点个数
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}
