package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1 == nil {
			tail.Next = list2
			break
		}
		if list2 == nil {
			tail.Next = list1
			break
		}
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	return dummy.Next
}
