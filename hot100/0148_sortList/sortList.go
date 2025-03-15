package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}
	half := slow.Next
	slow.Next = nil
	h1 := sortList(head)
	h2 := sortList(half)

	return combine(h1, h2)
}

func combine(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.Val < h2.Val {
		h1.Next = combine(h1.Next, h2)
		return h1
	} else {
		h2.Next = combine(h1, h2.Next)
		return h2
	}
}
