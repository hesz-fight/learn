package main

type ListNode struct {
	val  int
	next *ListNode
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	sum := 0
	for l1 != nil && l2 != nil {
		dummy.next = &ListNode{}
		if l1 == nil {
			sum = (l2.val + sum) / 10
			dummy.next.val = (l2.val + sum) % 10
			l2 = l2.next
		} else if l2 == nil {
			sum = (l1.val + sum) / 10
			dummy.next.val = (l1.val + sum) % 10
			l1 = l1.next
		} else {
			sum = (l1.val + l2.val + sum) / 10
			dummy.next.val = (l1.val + l2.val + sum) % 10
			l1 = l1.next
			l2 = l2.next
		}
		dummy = dummy.next
	}
	if sum != 0 {
		dummy.next = &ListNode{
			val: sum,
		}
	}

	return dummy.next
}
