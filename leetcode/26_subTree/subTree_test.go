package subtree

import "testing"

func Test_subtree(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				root1: root1,
				root2: root2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtree(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("subtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
