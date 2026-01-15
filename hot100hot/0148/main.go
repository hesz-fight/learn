package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治
func sortList(head *ListNode) *ListNode {
	// 对半拆分链表
	dummy := &ListNode{Next: head}

	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	half := slow.Next
	slow.Next = nil
	h1 := sortList(head)
	h2 := sortList(half)

	return combineList(h1, h2)
}

// 合并两个有序链表
func combineList(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.Val < h2.Val {
		h1.Next = combineList(h1.Next, h2)
		return h1
	} else {
		h2.Next = combineList(h1, h2.Next)
		return h2
	}
}
