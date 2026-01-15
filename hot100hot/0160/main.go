package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ta := headA
	tb := headB
	for headA != headB {
		if headA == nil {
			headA = tb
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = ta
		} else {
			headB = headB.Next
		}
	}

	return headA
}
