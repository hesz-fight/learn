package iimissingnum

import "testing"

func Test_fineMissingNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				nums: []int{0, 1, 3},
			},
			want: 2,
		},
		{
			name: "test",
			args: args{
				nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fineMissingNum(tt.args.nums); got != tt.want {
				t.Errorf("fineMissingNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
