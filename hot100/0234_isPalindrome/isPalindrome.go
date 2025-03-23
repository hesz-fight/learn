package ispalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

// 找到链表中点 然后反转前半部分 双指针比较
func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	half := slow.Next
	slow.Next = nil
	// 翻转后半部分
	if half != nil &&
		half.Next != nil {
		pre, cur := half, half.Next
		for cur != nil {
			next := cur.Next
			cur.Next = half
			half = cur
			pre.Next = next
			cur = next
		}
	}
	// 双指针比较
	for p1, p2 := head, half; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}

	return true
}

func revertList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pre := dummy.Next; pre.Next != nil; pre = pre.Next {
		next := pre.Next.Next      // 存储下一个处理的节点
		pre.Next.Next = dummy.Next // 指向头节点
		dummy.Next = pre.Next      // 更新头节点
		pre.Next = next            // 更新pre节点
	}
	return dummy.Next
}
