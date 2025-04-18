package findkthlargest

import "testing"

func Test_findKthLargest2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				nums: []int{2, 1},
				k:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
