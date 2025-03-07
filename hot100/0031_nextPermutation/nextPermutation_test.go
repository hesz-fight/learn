package nextpermutation

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{4, 5, 2, 6, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.nums)
			nextPermutation(tt.args.nums)
			fmt.Println(tt.args.nums)
		})

	}
}
