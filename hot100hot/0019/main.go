package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// [1,2,3,4,5]
// n = 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// find pre node
	p := dummy
	for i := 0; i < n; i++ {
		p = p.Next
	}
	t := dummy
	if p.Next != nil {
		p = p.Next
		t = t.Next
	}
	// t point to pre node
	t.Next = t.Next.Next
	return dummy.Next
}
