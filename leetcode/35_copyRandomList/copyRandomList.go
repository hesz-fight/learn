package copyrandomlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nHead := &Node{}
	var p, q *Node
	// 处理next指针域
	for p = head; p != nil; p = p.Next {
		nNode := &Node{Val: p.Val}
		if p == head {
			nHead = nNode
			q = nNode
			continue
		}
		q.Next = nNode
		q = q.Next
	}

	for p1, q1 := head, nHead; p1 != nil && q1 != nil; p1, q1 = p1.Next, q1.Next {
		// 外层循环找到不为nil的random指针
		if p1.Random == nil {
			continue
		}
		// 双指针找到random指针指向的节点
		for p2, q2 := head, nHead; p2 != nil && q2 != nil; p2, q2 = p2.Next, q2.Next {
			if p2 == p1.Random {
				q1.Random = q2
				break
			}
		}
	}

	return nHead
}

// 使用 map 存储老节点和新节点的映射 减少遍历次数
func copyRandomList2(head *Node) *Node {
	old2new := make(map[*Node]*Node)
	for p := head; p != nil; p = p.Next {
		old2new[p] = &Node{Val: p.Val}
	}

	for p := head; p != nil; p = p.Next {
		old2new[p].Next = old2new[p.Next]
		old2new[p].Random = old2new[p.Random]
	}

	return old2new[head]
}
