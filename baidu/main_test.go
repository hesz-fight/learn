package baidu

import (
	"testing"
)

func Test_deleteNNode(t *testing.T) {
	type args struct {
		head *Node
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				head: &Node{
					Val: 1,
					Next: &Node{
						Val: 3,
						Next: &Node{
							Val: 5,
						},
					},
				},
				n: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNNode(tt.args.head, tt.args.n)
			printListNode(t, got)
		})
	}
}

func printListNode(t *testing.T, head *Node) {
	for p := head; p != nil; p = p.Next {
		t.Log(p.Val)
	}
}
