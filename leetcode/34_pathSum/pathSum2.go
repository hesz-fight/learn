package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tag int
var ans [][]int

func pathSum(root *TreeNode, target int) [][]int {
	tag = target
	ans = make([][]int, 0, 64)
	backRecur(root, []int{}, 0)

	return ans
}

func backRecur(root *TreeNode, path []int, pathSum int) {
	if root == nil {
		return
	}
	pathSum += root.Val
	path = append(path, root.Val)
	if root.Left == nil &&
		root.Right == nil &&
		pathSum == tag {
		ans = append(ans, append([]int{}, path...))
	}
	backRecur(root.Left, path, pathSum)
	backRecur(root.Right, path, pathSum)
}
