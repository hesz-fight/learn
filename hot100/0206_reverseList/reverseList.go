package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		next := cur.Next
		cur.Next = head
		pre.Next = next
		head = cur
		cur = next
	}

	return head
}

// 1 2
func revert(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	// 头插法反转 pre.Next
	for pre := dummy.Next; pre != nil && pre.Next != nil; {
		tmp := pre.Next.Next
		pre.Next.Next = dummy.Next
		dummy.Next = pre.Next
		pre.Next = tmp
		pre = tmp
	}

	return dummy.Next
}
