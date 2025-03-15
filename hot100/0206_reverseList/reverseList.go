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
