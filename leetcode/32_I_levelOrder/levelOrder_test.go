package ilevelorder

import (
	"reflect"
	"testing"
)

func Test_ilevelorder(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				root: root,
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ilevelorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ilevelorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
