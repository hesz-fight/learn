package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0 // 进位
	head := &ListNode{}
	tail := head
	for l1 != nil || l2 != nil {
		tail.Next = &ListNode{}
		tail = tail.Next
		if l1 == nil {
			tail.Val = (l2.Val + sum) % 10
			sum = (l2.Val + sum) / 10
			l2 = l2.Next
		} else if l2 == nil {
			tail.Val = (l1.Val + sum) % 10
			sum = (l1.Val + sum) / 10
			l1 = l1.Next
		} else {
			tail.Val = (l1.Val + l2.Val + sum) % 10
			sum = (l1.Val + l2.Val + sum) / 10
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if sum != 0 {
		tail.Next = &ListNode{Val: sum}
	}

	return head.Next
}
