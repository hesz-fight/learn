package firstcommonnode

type Node struct {
	Val  int
	Next *Node
}

func firstCommonnode(head1 *Node, head2 *Node) *Node {
	for head1 != nil && head2 != nil {
		if head1.Val == head2.Val {
			return head1
		}
		head1 = head1.Next
		if head1 == nil {
			head1 = head2
		}
		head2 = head2.Next
		if head2 == nil {
			head1 = head2
		}
	}

	return nil
}
