package baidu

type Node struct {
	Val  int
	Next *Node
}

// 1->3->5 n=1
// 删除一个单向链表的倒数第n个节点
func deleteNNode(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	p1, p2 := dummy, dummy
	for i := 1; i <= n; i++ {
		p1 = p1.Next
	}
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}
