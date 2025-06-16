package buildtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mp = map[int]int{}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for ind, elem := range inorder {
		mp[elem] = ind
	}
	return helper(preorder, inorder)
}

func helper(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	// 找到中序遍历中根节点的下标
	ind := getIndxOfInorder(preorder[0])
	root.Left = helper(preorder[1:ind+1], inorder[:ind])
	root.Right = helper(preorder[ind+1:], inorder[ind+1:])
	return root
}

func getIndxOfInorder(val int) int {
	return mp[val]
}
