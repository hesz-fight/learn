package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumny := &ListNode{Next: head}
	p1, p2 := dumny, dumny
	for i := 1; i <= n; i++ {
		if p1 != nil {
			p1 = p1.Next
		}
	}

	for p1 != nil && p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dumny.Next
}
