package generateparenthesis

var (
	ans []string
	tmp []byte
)

func generateParenthesis(n int) []string {
	{
		ans = []string{}
		tmp = []byte{}
	}
	dfs(0, 0, n)
	return ans
}

func dfs(left int, right int, n int) {
	if left+right == 2*n {
		ans = append(ans, string(tmp))
		return
	}
	if left >= right {
		if left < n {
			tmp = append(tmp, '(')
			dfs(left+1, right, n)
			tmp = tmp[:len(tmp)-1]
		}
		if left != 0 && right < n {
			tmp = append(tmp, ')')
			dfs(left, right+1, n)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
