package detectcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	flag := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			flag = true
			break
		}
	}
	if !flag {
		return nil
	}
	// head 和 slow 一起往前走
	// 两个指针会在入口处相遇
	for slow != head {
		head = head.Next
		slow = slow.Next
	}

	return head
}
