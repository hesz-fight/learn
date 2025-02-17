package treetodoublylist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var head *TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var pre *TreeNode
	inorder(root, pre)
	head.Left = pre
	pre.Right = head

	return head
}

func inorder(root *TreeNode, pre *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, pre)
	if pre == nil {
		head = root
	} else {
		pre.Right = root
		root.Left = pre
	}
	pre = root
	inorder(root.Right, pre)
}

// 栈实现中序遍历
func treeToDoublyList2(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	var head, pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre == nil {
			head = root
		} else {
			pre.Right = root
			root.Left = pre
		}
		pre = root
		// root节点被处理完毕前左子树也被处理完了 使root节点指向右子树
		root = root.Right
	}
	head.Left = pre
	pre.Right = head

	return head
}
