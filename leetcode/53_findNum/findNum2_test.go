package findnum

import "testing"

func Test_findNum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{4, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNum2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
